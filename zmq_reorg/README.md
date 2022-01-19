# ZMQ re-org test

This playground tests what bitcoind sends along its ZMQ in a re-org situation. 

## What this repo contains:

- `regtest.yml` & `regtest.sh`: These are used to spin up and control 2 regtest
  bitcoind nodes (btc1 & btc2). Note that initially, the nodes are _not_
connected.
- `client.go`: A go main file which spins up a client that connects to one of the bitcoind nodes. The client has two
  different modes it can be run in: `btcclient` and `zmq`. With `zmq`, it will print out exactly what is returned from a
  zmq subscription to bitcoind. With `btcclient`, it crealtes a btcwallet `BitcoindConn` and uses that to subscribe to
  notifications and the prints out the notifications it gets. It always uses the btc1 node.

## How to setup the test:

1. Ensure docker is running
2. While in this dir, run `./regtest.sh start` to spin up the 2 bitcoin nodes.
3. Start the client: zmq mode: `go run ./client zmq`, or btcclient mode: `go run ./client btcclient`

Ok now you can play around with the following commands to see what the client
receives. 

- `./regtest.sh mine1 3`: let node btc1 mine 3 blocks (defaults is 6 blocks).
  Change to `./regtest mine2` to mine blocks on btc2 rather.
- `./regtest.sh connect`: this will connect the two bitcoin nodes.
- `./regtest.sh diconnect`: disconnect the two bitcoin nodes. 

## Simulate a re-org:
1. Connect the bitcoin nodes
2. Mine some blocks on one of them (doesnt matter which one). Should see that
   the client gets all of them)
3. Disconnect the nodes.
4. Generate a few blocks on btc1
5. Generate a few blocks on btc2. Ensure that you mine more blocks on btc2 than
   you did on btc1.
6. Reconnect the nodes. Since the chain on btc2 is longer, this should cause
   btc1, and hence the client, to go through a re-org.

