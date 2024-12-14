package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) RegisterUser(ctx contractapi.TransactionContextInterface, userID string, userType string, realInfoHash string) error {
	user := User{
		UserID:       userID,
		UserType:     userType,
		RealInfoHash: realInfoHash,
		ProductList:  []*Product{},
	}
	userAsBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(userID, userAsBytes)
	if err != nil {
		return err
	}
	return nil
}

func (s *SmartContract) Uplink(ctx contractapi.TransactionContextInterface, userID string, traceabilityCode string, arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 string, arg7 string) (string, error) {
	userType, err := s.GetUserType(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("failed to get user type: %v", err)
	}
	ProductAsBytes, err := ctx.GetStub().GetState(traceabilityCode)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	var product Product
	if ProductAsBytes != nil {
		err = json.Unmarshal(ProductAsBytes, &product)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal product: %v", err)
		}
	}
	txtime, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return "", fmt.Errorf("failed to read TxTimestamp: %v", err)
	}
	timeLocation, _ := time.LoadLocation("Asia/Shanghai")
	format := time.Unix(txtime.Seconds, 0).In(timeLocation).Format("2006-01-02 15:04:05")
	txid := ctx.GetStub().GetTxID()
	product.TraceabilityCode = traceabilityCode
	switch userType {
	case "养蜂场":
		product.BeeFarmInput.BeeFarmName = arg1
		product.BeeFarmInput.BeeFarmLocation = arg2
		product.BeeFarmInput.BeeBoxId = arg3
		product.BeeFarmInput.HoneyVariety = arg4
		product.BeeFarmInput.FlowerVariety = arg5
		product.BeeFarmInput.Fa_IPFSCID = arg6
		product.BeeFarmInput.Fa_IPFSFileName = arg7
		product.BeeFarmInput.BeeFarmTxid = txid
		product.BeeFarmInput.BeeFarmTimestamp = format
	case "加工厂":
		product.ProcessingPlantInput.ProcessingPlantName = arg1
		product.ProcessingPlantInput.ProcessingPlantLocation = arg2
		product.ProcessingPlantInput.ProcessingBatchId = arg3
		product.ProcessingPlantInput.PackagingSpecification = arg4
		product.ProcessingPlantInput.ShelfLife = arg5
		product.ProcessingPlantInput.Fac_IPFSCID = arg6
		product.ProcessingPlantInput.Fac_IPFSFileName = arg7
		product.ProcessingPlantInput.ProcessingPlantTxid = txid
		product.ProcessingPlantInput.ProcessingPlantTimestamp = format
	case "批发商":
		product.WholesalerInput.WarehouseName = arg1
		product.WholesalerInput.WarehouseLocation = arg2
		product.WholesalerInput.WholesalerBatchId = arg3
		product.WholesalerInput.TransportationMethod = arg4
		product.WholesalerInput.TransportMode = arg5
		product.WholesalerInput.Dr_IPFSCID = arg6
		product.WholesalerInput.Dr_IPFSFileName = arg7
		product.WholesalerInput.WholesalerTxid = txid
		product.WholesalerInput.WholesalerTimestamp = format
	case "零售商":
		product.RetailerInput.StoreName = arg1
		product.RetailerInput.StoreLocation = arg2
		product.RetailerInput.RetailerBatchId = arg3
		product.RetailerInput.SalesChannel = arg4
		product.RetailerInput.SalesPrice = arg5
		product.RetailerInput.Sh_IPFSCID = arg6
		product.RetailerInput.Sh_IPFSFileName = arg7
		product.RetailerInput.RetailerTxid = txid
		product.RetailerInput.RetailerTimestamp = format
	}
	productAsBytes, err := json.Marshal(product)
	if err != nil {
		return "", fmt.Errorf("failed to marshal product: %v", err)
	}
	err = ctx.GetStub().PutState(traceabilityCode, productAsBytes)
	if err != nil {
		return "", fmt.Errorf("failed to put product: %v", err)
	}
	err = s.AddProduct(ctx, userID, &product)
	if err != nil {
		return "", fmt.Errorf("failed to add product to user: %v", err)
	}
	return txid, nil
}

func (s *SmartContract) AddProduct(ctx contractapi.TransactionContextInterface, userID string, product *Product) error {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return fmt.Errorf("the user %s does not exist", userID)
	}
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return err
	}
	user.ProductList = append(user.ProductList, product)
	userAsBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(userID, userAsBytes)
	if err != nil {
		return err
	}
	return nil
}

func (s *SmartContract) GetUserType(ctx contractapi.TransactionContextInterface, userID string) (string, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return "", fmt.Errorf("the user %s does not exist", userID)
	}
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return "", err
	}
	return user.UserType, nil
}

func (s *SmartContract) GetUserInfo(ctx contractapi.TransactionContextInterface, userID string) (*User, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return &User{}, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return &User{}, fmt.Errorf("the user %s does not exist", userID)
	}
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

func (s *SmartContract) GetProductInfo(ctx contractapi.TransactionContextInterface, traceabilityCode string) (*Product, error) {
	ProductAsBytes, err := ctx.GetStub().GetState(traceabilityCode)
	if err != nil {
		return &Product{}, fmt.Errorf("failed to read from world state: %v", err)
	}
	var product Product
	if ProductAsBytes != nil {
		err = json.Unmarshal(ProductAsBytes, &product)
		if err != nil {
			return &Product{}, fmt.Errorf("failed to unmarshal product: %v", err)
		}
	}
	return &product, nil
}

func (s *SmartContract) GetProductList(ctx contractapi.TransactionContextInterface, userID string) ([]*Product, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return nil, fmt.Errorf("the user %s does not exist", userID)
	}
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, err
	}
	return user.ProductList, nil
}

func (s *SmartContract) GetAllProductInfo(ctx contractapi.TransactionContextInterface) ([]Product, error) {
	productListAsBytes, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	defer func(productListAsBytes shim.StateQueryIteratorInterface) {
		_ = productListAsBytes.Close()
	}(productListAsBytes)
	var products []Product
	for productListAsBytes.HasNext() {
		queryResponse, err := productListAsBytes.Next()
		if err != nil {
			return nil, err
		}
		var product Product
		err = json.Unmarshal(queryResponse.Value, &product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (s *SmartContract) GetProductHistory(ctx contractapi.TransactionContextInterface, traceabilityCode string) ([]HistoryQueryResult, error) {
	log.Printf("Get Asset History: ID %v", traceabilityCode)
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(traceabilityCode)
	if err != nil {
		return nil, err
	}
	defer func(resultsIterator shim.HistoryQueryIteratorInterface) {
		_ = resultsIterator.Close()
	}(resultsIterator)
	var records []HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var product Product
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &product)
			if err != nil {
				return nil, err
			}
		} else {
			product = Product{
				TraceabilityCode: traceabilityCode,
			}
		}
		//goland:noinspection GoDeprecation
		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}
		targetLocation, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			return nil, err
		}
		timestamp = timestamp.In(targetLocation)
		formattedTime := timestamp.Format("2006-01-02 15:04:05")
		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: formattedTime,
			Record:    &product,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}
	return records, nil
}
