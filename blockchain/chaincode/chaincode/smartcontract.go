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
		FruitList:    []*Fruit{},
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

func (s *SmartContract) Uplink(ctx contractapi.TransactionContextInterface, userID string, traceability_code string, arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 string, arg7 string) (string, error) {
	userType, err := s.GetUserType(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("Failed to get user type: %v", err)
	}
	FruitAsBytes, err := ctx.GetStub().GetState(traceability_code)
	if err != nil {
		return "", fmt.Errorf("Failed to read from world state: %v", err)
	}
	var fruit Fruit
	if FruitAsBytes != nil {
		err = json.Unmarshal(FruitAsBytes, &fruit)
		if err != nil {
			return "", fmt.Errorf("Failed to unmarshal fruit: %v", err)
		}
	}
	txtime, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return "", fmt.Errorf("Failed to read TxTimestamp: %v", err)
	}
	timeLocation, _ := time.LoadLocation("Asia/Shanghai")
	time := time.Unix(txtime.Seconds, 0).In(timeLocation).Format("2006-01-02 15:04:05")
	txid := ctx.GetStub().GetTxID()
	fruit.Traceability_code = traceability_code
	switch userType {
	case "种植户":
		fruit.Farmer_input.Fa_fruitName = arg1
		fruit.Farmer_input.Fa_origin = arg2
		fruit.Farmer_input.Fa_plantTime = arg3
		fruit.Farmer_input.Fa_pickingTime = arg4
		fruit.Farmer_input.Fa_farmerName = arg5
		fruit.Farmer_input.Fa_IPFSCID = arg6
		fruit.Farmer_input.Fa_IPFSFileName = arg7
		fruit.Farmer_input.Fa_Txid = txid
		fruit.Farmer_input.Fa_Timestamp = time
	case "工厂":
		fruit.Factory_input.Fac_productName = arg1
		fruit.Factory_input.Fac_productionbatch = arg2
		fruit.Factory_input.Fac_productionTime = arg3
		fruit.Factory_input.Fac_factoryName = arg4
		fruit.Factory_input.Fac_contactNumber = arg5
		fruit.Factory_input.Fac_IPFSCID = arg6
		fruit.Factory_input.Fac_IPFSFileName = arg7
		fruit.Factory_input.Fac_Txid = txid
		fruit.Factory_input.Fac_Timestamp = time
	case "运输司机":
		fruit.Driver_input.Dr_name = arg1
		fruit.Driver_input.Dr_age = arg2
		fruit.Driver_input.Dr_phone = arg3
		fruit.Driver_input.Dr_carNumber = arg4
		fruit.Driver_input.Dr_transport = arg5
		fruit.Driver_input.Dr_IPFSCID = arg6
		fruit.Driver_input.Dr_IPFSFileName = arg7
		fruit.Driver_input.Dr_Txid = txid
		fruit.Driver_input.Dr_Timestamp = time
	case "商店":
		fruit.Shop_input.Sh_storeTime = arg1
		fruit.Shop_input.Sh_sellTime = arg2
		fruit.Shop_input.Sh_shopName = arg3
		fruit.Shop_input.Sh_shopAddress = arg4
		fruit.Shop_input.Sh_shopPhone = arg5
		fruit.Shop_input.Sh_IPFSCID = arg6
		fruit.Shop_input.Sh_IPFSFileName = arg7
		fruit.Shop_input.Sh_Txid = txid
		fruit.Shop_input.Sh_Timestamp = time
	}
	fruitAsBytes, err := json.Marshal(fruit)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal fruit: %v", err)
	}
	err = ctx.GetStub().PutState(traceability_code, fruitAsBytes)
	if err != nil {
		return "", fmt.Errorf("Failed to put fruit: %v", err)
	}
	err = s.AddFruit(ctx, userID, &fruit)
	if err != nil {
		return "", fmt.Errorf("Failed to add fruit to user: %v", err)

	}
	return txid, nil
}

func (s *SmartContract) AddFruit(ctx contractapi.TransactionContextInterface, userID string, fruit *Fruit) error {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return fmt.Errorf("Failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return fmt.Errorf("The user %s does not exist", userID)
	}
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return err
	}
	user.FruitList = append(user.FruitList, fruit)
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
		return "", fmt.Errorf("Failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return "", fmt.Errorf("The user %s does not exist", userID)
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
		return &User{}, fmt.Errorf("Failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return &User{}, fmt.Errorf("The user %s does not exist", userID)
	}
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

func (s *SmartContract) GetFruitInfo(ctx contractapi.TransactionContextInterface, traceability_code string) (*Fruit, error) {
	FruitAsBytes, err := ctx.GetStub().GetState(traceability_code)
	if err != nil {
		return &Fruit{}, fmt.Errorf("Failed to read from world state: %v", err)
	}
	var fruit Fruit
	if FruitAsBytes != nil {
		err = json.Unmarshal(FruitAsBytes, &fruit)
		if err != nil {
			return &Fruit{}, fmt.Errorf("Failed to unmarshal fruit: %v", err)
		}
	}
	return &fruit, nil
}

func (s *SmartContract) GetFruitList(ctx contractapi.TransactionContextInterface, userID string) ([]*Fruit, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return nil, fmt.Errorf("The user %s does not exist", userID)
	}
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, err
	}
	return user.FruitList, nil
}

func (s *SmartContract) GetAllFruitInfo(ctx contractapi.TransactionContextInterface) ([]Fruit, error) {
	fruitListAsBytes, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state: %v", err)
	}
	defer func(fruitListAsBytes shim.StateQueryIteratorInterface) {
		_ = fruitListAsBytes.Close()
	}(fruitListAsBytes)
	var fruits []Fruit
	for fruitListAsBytes.HasNext() {
		queryResponse, err := fruitListAsBytes.Next()
		if err != nil {
			return nil, err
		}
		var fruit Fruit
		err = json.Unmarshal(queryResponse.Value, &fruit)
		if err != nil {
			return nil, err
		}
		fruits = append(fruits, fruit)
	}
	return fruits, nil
}

func (s *SmartContract) GetFruitHistory(ctx contractapi.TransactionContextInterface, traceability_code string) ([]HistoryQueryResult, error) {
	log.Printf("Get Asset History: ID %v", traceability_code)
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(traceability_code)
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
		var fruit Fruit
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &fruit)
			if err != nil {
				return nil, err
			}
		} else {
			fruit = Fruit{
				Traceability_code: traceability_code,
			}
		}
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
			Record:    &fruit,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}
	return records, nil
}
