#!/bin/bash

echo "Starting the blockchain network..."

./stop.sh

docker run --name fabrictrace-mysql -p 3337:3306 -e MYSQL_ROOT_PASSWORD=fabrictrace -d mysql:8

image_versions=("2.5.6")
images=("hyperledger/fabric-tools" "hyperledger/fabric-peer" "hyperledger/fabric-orderer" "hyperledger/fabric-ccenv" "hyperledger/fabric-baseos")

for image in "${images[@]}"; do
    for version in "${image_versions[@]}"; do
        if ! docker images -a | grep "$image" | grep "$version" &> /dev/null; then
            echo "Image $image:$version does not exist, starting to pull..."
            docker pull "$image:$version"
        fi
    done
done

./network.sh up createChannel
./network.sh deployCC -ccn trace -ccp ../chaincode -ccl go

cp -r organizations explorer/

docker compose -f explorer/docker-compose.yaml up -d

echo "The blockchain network is now running."
