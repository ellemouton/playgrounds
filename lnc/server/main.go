package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"

	"github.com/lightninglabs/lightning-node-connect/gbn"
	"github.com/lightningnetwork/lnd"
	"github.com/lightningnetwork/lnd/build"
	"github.com/lightningnetwork/lnd/signal"

	"github.com/lightninglabs/lightning-node-connect/itest/mockrpc"

	"github.com/btcsuite/btcd/btcec"
	"github.com/lightninglabs/lightning-node-connect/mailbox"
	"github.com/lightningnetwork/lnd/keychain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	insecure      = flag.Bool("insecure", false, "")
	serverAddr    = flag.String("serveraddr", "mailbox.terminal.lightning.today:443", "")
	pairingSecret = flag.String("pw", "", "")
	reusePassword = flag.Bool("reuse", false, "")
	authData      = flag.String("authData", "", "")
)

func main() {
	flag.Parse()

	if err := startServer(); err != nil {
		log.Fatalln(err)
	}
}

func startServer() error {
	// Set up dem logs.
	logWriter := build.NewRotatingLogWriter()
	interceptor, _ := signal.Intercept()
	lnd.AddSubLogger(logWriter, gbn.Subsystem, interceptor, gbn.UseLogger)
	lnd.AddSubLogger(logWriter, mailbox.Subsystem, interceptor, mailbox.UseLogger)
	logWriter.SetLogLevels("trace")

	pwStr, password, err := getPassword()
	if err != nil {
		return err
	}
	fmt.Println(password)

	if err := storePassword(pwStr); err != nil {
		return err
	}

	// Start the mailbox gRPC server.
	fmt.Printf("Connecting to %s, tls? %v\n", *serverAddr, !*insecure)
	tlsConfig := &tls.Config{}
	if *insecure {
		tlsConfig = &tls.Config{InsecureSkipVerify: true}
	}
	mailboxServer, err := mailbox.NewServer(
		*serverAddr, password[:],
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	)
	if err != nil {
		return err
	}

	privateKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return err
	}

	ecdh := &keychain.PrivKeyECDH{PrivKey: privateKey}
	noiseConn := mailbox.NewNoiseGrpcConn(
		ecdh, []byte(*authData), password[:],
	)

	server := grpc.NewServer(grpc.Creds(noiseConn))

	mockrpc.RegisterMockServiceServer(server, &mockrpc.Server{})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		s := <-interceptor.ShutdownChannel()
		log.Printf("got signal %v, attempting graceful shutdown", s)
		server.GracefulStop()
	}()

	fmt.Printf("Mock RPC server listening on %s\n", mailboxServer.Addr())

	log.Println("starting grpc server")
	err = server.Serve(mailboxServer)
	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}

	wg.Wait()
	log.Println("clean shutdown")

	return nil
}

func getPassword() (string, []byte, error) {
	if *reusePassword {
		data, err := ioutil.ReadFile("password.txt")
		if err != nil {
			return "", nil, err
		}

		fmt.Printf("Reusing previous password: %s\n", string(data))
		words := strings.Split(string(data), " ")
		var mnemonicWords [mailbox.NumPasswordWords]string
		copy(mnemonicWords[:], words)
		pw := mailbox.PasswordMnemonicToEntropy(mnemonicWords)
		return string(data), pw[:], nil
	}

	if *pairingSecret != "" {
		fmt.Printf("Using given password: %s\n", *pairingSecret)
		words := strings.Split(*pairingSecret, " ")
		var mnemonicWords [mailbox.NumPasswordWords]string
		copy(mnemonicWords[:], words)
		pw := mailbox.PasswordMnemonicToEntropy(mnemonicWords)
		return *pairingSecret, pw[:], nil
	}

	strs, password, err := mailbox.NewPassword()
	str := strings.Join(strs[:], " ")

	fmt.Printf("Generated new password: %s\n", str)

	return str, password[:], err
}

func storePassword(pw string) error {
	return ioutil.WriteFile("password.txt", []byte(pw), 0777)
}
