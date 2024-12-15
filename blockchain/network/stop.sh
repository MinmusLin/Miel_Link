echo "Shutting down the blockchain network..."

docker compose -f explorer/docker-compose.yaml down -v > /dev/null 2>&1

./network.sh down > /dev/null 2>&1

docker rm -f miellink-mysql > /dev/null 2>&1

docker compose -f ipfs.yaml down -v

rm -rf explorer/organizations
rm -rf ipfs
rm -rf ../../application/backend/files

mkdir -p ../../application/backend/files/downloadfiles
mkdir -p ../../application/backend/files/uploadfiles

echo "The blockchain network has been shut down."
