docker compose -f explorer/docker-compose.yaml down -v 
./network.sh down
rm -rf explorer/organizations
docker rm -f fabrictrace-mysql
docker compose -f ipfs.yaml down -v