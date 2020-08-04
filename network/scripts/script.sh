#!/bin/bash

DELAY="3"
TIMEOUT="10"
VERBOSE="false"
COUNTER=1
MAX_RETRY=5

CC_SRC_PATH="irscc/"

createChannel() {
	CORE_PEER_LOCALMSPID=partya
	CORE_PEER_ADDRESS=irs-partya:7051
	CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/partya.example.com/users/Admin@partya.example.com/msp
	echo "===================== Creating channel ===================== "
	peer channel create -o irs-orderer:7050 -c irs -f ./channel-artifacts/channel.tx
	echo "===================== Channel created ===================== "
}

joinChannel () {
	for org in partya partyb 
	do
		CORE_PEER_LOCALMSPID=$org
		CORE_PEER_ADDRESS=irs-$org:7051
		CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$org.example.com/users/Admin@$org.example.com/msp
		echo "===================== Org $org joining channel ===================== "
		peer channel join -b irs.block -o irs-orderer:7050
		echo "===================== Channel joined ===================== "
	done
}
updateAnchorPeer() {
	for org in partya partyb 
	do
	echo "====================> Configure this peer $org to be a Anchor Peer"
		CORE_PEER_LOCALMSPID=$org
		CORE_PEER_ADDRESS=irs-$org:7051
		CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$org.example.com/users/Admin@$org.example.com/msp
		peer channel update -o irs-orderer:7050 -c irs -f "./channel-artifacts/$org.tx"
    done
} 

installChaincode() {
	for org in partya partyb 
	do
		CORE_PEER_LOCALMSPID=$org
		CORE_PEER_ADDRESS=irs-$org:7051
		CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$org.example.com/users/Admin@$org.example.com/msp
		echo "===================== Org $org installing chaincode ===================== "
		peer chaincode install -n irscc -v 0 -l golang -p  ${CC_SRC_PATH}
		echo "===================== Org $org chaincode installed ===================== "
	done
}

instantiateChaincode() {
	CORE_PEER_LOCALMSPID=partya
	CORE_PEER_ADDRESS=irs-partya:7051
	CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/partya.example.com/users/Admin@partya.example.com/msp
	echo "===================== Instantiating chaincode ===================== "
	peer chaincode instantiate -o irs-orderer:7050 -C irs -n irscc -v 0 -c '{"Args":["init"]}' -P "OR('partya.peer','partyb.peer')"
	
	#peer chaincode instantiate -o irs-orderer:7050 -C irs -n irscc -v 0 -c '{"Args":["init"]}' 
	echo "===================== Chaincode instantiated ===================== "
}

## Create channel
sleep 1
echo "Creating channel..."
createChannel

## Join all the peers to the channel
echo "Having all peers join the channel..."
joinChannel

## Update anchorpeer
echo "Having all anchorpeer updated..."
updateAnchorPeer

## Install chaincode on all peers
echo "Installing chaincode..."
installChaincode

# Instantiate chaincode
echo "Instantiating chaincode..."
instantiateChaincode

echo
echo "========= IRS network sample setup completed =========== "
echo

exit 0