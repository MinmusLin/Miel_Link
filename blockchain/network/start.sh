#!/bin/bash

./stop.sh

image_versions=("2.5.6")

images=("hyperledger/fabric-tools" "hyperledger/fabric-peer" "hyperledger/fabric-orderer" "hyperledger/fabric-ccenv" "hyperledger/fabric-baseos")

for image in "${images[@]}"
do
    for version in "${image_versions[@]}"
    do
        if ! docker images -a | grep "$image" | grep "$version" &> /dev/null
        then
            echo "镜像 $image:$version 不存在，开始拉取..."
            docker pull "$image:$version"
        fi
    done
done

./network.sh up createChannel
./network.sh deployCC -ccn trace -ccp ../chaincode -ccl go
cp -r organizations explorer/
docker run --name fabrictrace-mysql -p 3337:3306 -e MYSQL_ROOT_PASSWORD=fabrictrace -d mysql:8
docker compose -f ipfs.yaml up -d
docker compose -f explorer/docker-compose.yaml up -d