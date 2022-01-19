#!/bin/bash

COMPOSE="docker-compose -f ./regtest.yml -p regtest"

function start() {
        $COMPOSE up --force-recreate -d

        echo "wait for nodes to start"
        sleep 5 
        
        btc1 createwallet booooop
        btc2 createwallet booooop2
}

function stop() {
        $COMPOSE down --volumes
}

function btc1() {
        docker exec -ti -u bitcoin regtest_btc1_1 bitcoin-cli -regtest -rpcuser=lightning -rpcpassword=lightning "$@"
}

function btc2() {
        docker exec -ti -u bitcoin regtest_btc2_1 bitcoin-cli -regtest -rpcuser=lightning -rpcpassword=lightning "$@"
}

function mine1() {
        NUMBLOCKS=6
        if [ ! -z "$1" ]
        then 
        NUMBLOCKS=$1
        fi

        btc1 -generate $NUMBLOCKS
}

function mine2() {
        NUMBLOCKS=6
        if [ ! -z "$1" ]
        then 
        NUMBLOCKS=$1
        fi

        btc2 -generate $NUMBLOCKS
}

function connect() {
        echo "connecting peer"
        btc2 addnode regtest_btc1_1:18444 onetry
}

function disconnect() {
        echo "disconnecting peer"
        btc2 disconnectnode regtest_btc1_1:18444
}

if [ $# -lt 1 ]
then 
  echo "Usage: $0 start|stop|mine1|mine2|connect|disconnect"
fi

CMD=$1
shift
$CMD "$@"
