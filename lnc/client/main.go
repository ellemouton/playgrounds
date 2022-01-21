package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"crypto/sha512"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/btcsuite/btcd/btcec"
	"github.com/lightninglabs/lightning-node-connect/gbn"
	"github.com/lightninglabs/lightning-node-connect/itest/mockrpc"
	"github.com/lightninglabs/lightning-node-connect/mailbox"
	"github.com/lightningnetwork/lnd"
	"github.com/lightningnetwork/lnd/build"
	"github.com/lightningnetwork/lnd/keychain"
	"github.com/lightningnetwork/lnd/signal"
	"google.golang.org/grpc"
)

var (
	insecure      = flag.Bool("insecure", false, "")
	serverAddr    = flag.String("serveraddr", "mailbox.terminal.lightning.today:443", "")
	pairingSecret = flag.String("pw", "", "")
)

func main() {
	logWriter := build.NewRotatingLogWriter()
	interceptor, _ := signal.Intercept()
	lnd.AddSubLogger(logWriter, gbn.Subsystem, interceptor, gbn.UseLogger)
	lnd.AddSubLogger(logWriter, mailbox.Subsystem, interceptor, mailbox.UseLogger)
	logWriter.SetLogLevels("trace")

	flag.Parse()

	client, err := lndConn()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	c := mockrpc.NewMockServiceClient(client)

	quit := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := chatWithServer(c, quit); err != nil {
			log.Fatal(err)
		}
	}()

	<-interceptor.ShutdownChannel()
	log.Printf("got signal. attempting graceful shutdown")
	close(quit)
	client.Close()
	wg.Wait()
}

func chatWithServer(c mockrpc.MockServiceClient, quit chan struct{}) error {
	buf := bufio.NewReader(os.Stdin)

	counter := 0
	ctx := context.Background()
	for {
		select {
		case <-quit:
			return nil
		default:
		}

		fmt.Println("> press enter to send things!")
		_, err := buf.ReadBytes('\n')
		if err != nil {
			return err
		}

		largeResp := make([]byte, 1024*4)
		rand.Read(largeResp)
		req := &mockrpc.Request{Req: largeResp}

		_, err = c.MockServiceMethod(ctx, req)
		if err != nil {
			return err
		}

		fmt.Printf("got the thing: %d\n", counter)
		counter++
	}

	return nil
}

func lndConn() (*grpc.ClientConn, error) {
	words := strings.Split(*pairingSecret, " ")
	var mnemonicWords [mailbox.NumPasswordWords]string
	copy(mnemonicWords[:], words)
	password := mailbox.PasswordMnemonicToEntropy(mnemonicWords)
	fmt.Println(password)

	sid := sha512.Sum512(password[:])
	receiveSID := mailbox.GetSID(sid, true)
	sendSID := mailbox.GetSID(sid, false)

	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}
	ecdh := &keychain.PrivKeyECDH{PrivKey: privKey}

	ctx := context.Background()
	transportConn := mailbox.NewClientConn(ctx, receiveSID, sendSID)
	noiseConn := mailbox.NewNoiseGrpcConn(ecdh, nil, password[:])

	dialOpts := []grpc.DialOption{
		grpc.WithContextDialer(transportConn.Dial),
		grpc.WithTransportCredentials(noiseConn),
		grpc.WithPerRPCCredentials(noiseConn),
	}

	tlsConfig := &tls.Config{}
	if *insecure {
		tlsConfig = &tls.Config{InsecureSkipVerify: true}
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = tlsConfig

	fmt.Printf("server addr: %s, tls? %v", *serverAddr, !*insecure)
	return grpc.DialContext(ctx, *serverAddr, dialOpts...)
}
