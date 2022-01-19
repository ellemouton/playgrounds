package main

import (
	"bytes"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcwallet/chain"
	"github.com/lightninglabs/gozmq"
	"github.com/lightningnetwork/lnd/signal"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

const (
	// rawBlockZMQCommand is the command used to receive raw block
	// notifications from bitcoind through ZMQ.
	rawBlockZMQCommand = "rawblock"

	// rawTxZMQCommand is the command used to receive raw transaction
	// notifications from bitcoind through ZMQ.
	rawTxZMQCommand = "rawtx"

	// seqNumLen is the length of the sequence number of a message sent from
	// bitcoind through ZMQ.
	seqNumLen = 4

	// maxRawBlockSize is the maximum size in bytes for a raw block received
	// from bitcoind through ZMQ.
	maxRawBlockSize = 4e6
)

func main() {
	interceptor, err := signal.Intercept()
	if err != nil {
		log.Fatalln(err)
	}

	subscription := "btcclient"
	if len(os.Args) > 1 {
		subscription = os.Args[1]
	}

	if subscription != "btcclient" && subscription != "zmq" {
		log.Fatalln("Must choose between `btcclient` or `zmq`. " +
			"Ex: start program with `go run ./client zmq`")
	}

	zmqBlockConn, err := gozmq.Subscribe(
		"localhost:28332", []string{rawBlockZMQCommand}, time.Second*5,
	)
	if err != nil {
		log.Fatalln(fmt.Errorf("unable to subscribe for zmq block "+
			"events: %v", err))
	}

	quit := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	switch subscription {
	case "btcclient":
		go func() {
			defer wg.Done()
			if err := btcClientNtfns(quit); err != nil {
				fmt.Println("cant do things", err)
			}
		}()

	case "zmq":
		go func() {
			defer wg.Done()
			if err := zmqOutput(zmqBlockConn, quit); err != nil {
				fmt.Println(err)
			}
		}()
	}

	<-interceptor.ShutdownChannel()
	fmt.Println("Received shutdown signal")

	close(quit)
	zmqBlockConn.Close()

	wg.Wait()
}

func zmqOutput(conn *gozmq.Conn, quit chan struct{}) error {
	var (
		command [len(rawBlockZMQCommand)]byte
		seqNum  [seqNumLen]byte
		data    = make([]byte, maxRawBlockSize)
	)

	fmt.Println("Waiting for blocks...")
	for {
		select {
		case <-quit:
			return nil
		default:
		}

		// Poll an event from the ZMQ socket.
		var (
			bufs = [][]byte{command[:], data, seqNum[:]}
			err  error
		)
		bufs, err = conn.Receive(bufs)
		if err != nil {
			// EOF should only be returned if the connection was
			// explicitly closed, so we can exit at this point.
			if err == io.EOF {
				return nil
			}

			return fmt.Errorf("unable to receive ZMQ %v "+
				"message: %v", rawBlockZMQCommand, err)
		}

		// We have an event! We'll now ensure it is a block event,
		// deserialize it, and report it to the different rescan
		// clients.
		eventType := string(bufs[0])
		switch eventType {
		case rawBlockZMQCommand:
			block := &wire.MsgBlock{}
			r := bytes.NewReader(bufs[1])
			if err := block.Deserialize(r); err != nil {
				return fmt.Errorf("unable to deserialize "+
					"block: %v", err)
			}

			fmt.Println("---------------BLOCK--------------------")
			fmt.Println(block.BlockHash())
			fmt.Printf("%+v\n", block)
			fmt.Println("----------------------------------------")
		default:
			return fmt.Errorf("received unexpected event type "+
				"from %v subscription: %v", rawBlockZMQCommand,
				eventType)
		}
	}

	return nil
}

func btcClientNtfns(quit chan struct{}) error {
	btc, err := chain.NewBitcoindConn(&chain.BitcoindConfig{
		ChainParams:     &chaincfg.RegressionNetParams,
		Host:            "localhost:18443",
		User:            "lightning",
		Pass:            "lightning",
		ZMQBlockHost:    "localhost:28332",
		ZMQTxHost:       "localhost:28333",
		ZMQReadDeadline: time.Second * 5,
		Dialer: func(addr string) (net.Conn, error) {
			return net.DialTimeout("tcp", addr, time.Second*5)
		},
	})
	if err != nil {
		return err
	}

	if err := btc.Start(); err != nil {
		return err
	}
	defer btc.Stop()

	client := btc.NewBitcoindClient()
	if err := client.Start(); err != nil {
		return err
	}
	defer client.Stop()

	if err := client.NotifyBlocks(); err != nil {
		return err
	}

	ntfns := client.Notifications()

	for {
		select {
		case thing := <-ntfns:
			switch thing.(type) {
			case chain.BlockConnected:
				block := thing.(chain.BlockConnected)
				fmt.Printf("Block Connected: height: %d, "+
					"hash: %s\n", block.Height, block.Hash)

			case chain.BlockDisconnected:
				block := thing.(chain.BlockDisconnected)
				fmt.Printf("Block Disconnected: height: %d, "+
					"hash: %s\n", block.Height, block.Hash)

			default:
				fmt.Printf("Got notification of type: %T\n",
					thing)
			}
		case <-quit:
			return nil
		}
	}
}
