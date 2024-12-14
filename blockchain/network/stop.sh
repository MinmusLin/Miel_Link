echo "Shutting down the blockchain network..."

docker compose -f explorer/docker-compose.yaml down -v > /dev/null 2>&1

./network.sh down > /dev/null 2>&1

rm -rf explorer/organizations

sudo rm -rf ipfs

sudo rm -rf ../../application/backend/files

docker rm -f miellink-mysql > /dev/null 2>&1

docker compose -f ipfs.yaml down -v

echo "The blockchain network has been shut down."
