package pkg

import (
	"crypto/x509"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	mspID        = "Org1MSP"
	cryptoPath   = "../../blockchain/network/organizations/peerOrganizations/org1.example.com"
	certPath     = cryptoPath + "/users/User1@org1.example.com/msp/signcerts/User1@org1.example.com-cert.pem"
	keyPath      = cryptoPath + "/users/User1@org1.example.com/msp/keystore/"
	tlsCertPath  = cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
	peerEndpoint = "localhost:7051"
	gatewayPeer  = "peer0.org1.example.com"
)

func ChaincodeQuery(fcn string, arg string) (string, error) {
	contract, conn, gw := GetContract()
	defer conn.Close()
	defer func(gw *client.Gateway) {
		_ = gw.Close()
	}(gw)
	evaluateResult, err := contract.EvaluateTransaction(fcn, arg)
	if err != nil {
		return "", fmt.Errorf("failed to evaluate transaction: %w", err)
	}
	fmt.Printf("Result: %s\n", evaluateResult)
	return string(evaluateResult), nil
}

func ChaincodeInvoke(fcn string, args []string) (string, error) {
	contract, conn, gw := GetContract()
	defer conn.Close()
	defer func(gw *client.Gateway) {
		_ = gw.Close()
	}(gw)
	submitResult, commit, err := contract.SubmitAsync(fcn, client.WithArguments(args...))
	if err != nil {
		return "", fmt.Errorf("failed to submit transaction asynchronously: %w", err)
	}
	fmt.Printf("\nSuccessfully submitted transaction: %v", string(submitResult))
	if commitStatus, err := commit.Status(); err != nil {
		return "", fmt.Errorf("failed to get commit status: %w", err)
	} else if !commitStatus.Successful {
		return "", fmt.Errorf("transaction %s failed to commit with status: %d", commitStatus.TransactionID, int32(commitStatus.Code))
	}
	fmt.Printf("Transaction committed successfully\n")
	return commit.TransactionID(), nil
}

func GetContract() (*client.Contract, *grpc.ClientConn, *client.Gateway) {
	clientConnection := newGrpcConnection()
	id := newIdentity()
	sign := newSign()
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	chaincodeName := "trace"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}
	channelName := "mychannel"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)
	return contract, clientConnection, gw
}

func newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(certPath)
	if err != nil {
		panic(err)
	}
	id, err := identity.NewX509Identity(mspID, certificate)
	if err != nil {
		panic(err)
	}
	return id
}

func newSign() identity.Sign {
	files, err := os.ReadDir(keyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := os.ReadFile(path.Join(keyPath, files[0].Name()))
	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}
	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}
	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}
	return sign
}

func newGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(tlsCertPath)
	if err != nil {
		panic(err)
	}
	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)
	connection, err := grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}
	return connection
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}
