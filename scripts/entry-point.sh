#!/bin/sh

CHAINID=$1
GENACCT=$2

if [ -z "$1" ]; then
  echo "Need to input chain id..."
  exit 1
fi

if [ -z "$2" ]; then
  echo "Need to input genesis account name..."
  exit 1
fi

coins="1000000000000000atn,1000000000000000mdm"
autonomy init --chain-id $CHAINID $CHAINID
autonomy keys add validator
autonomy add-genesis-account validator $coins
autonomy add-genesis-account $GENACCT $coins
autonomy gentx validator 100000000atn --chain-id $CHAINID
autonomy collect-gentxs

#autonomy start 
