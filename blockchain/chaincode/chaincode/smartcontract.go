package chaincode

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// 定义合约结构体
type SmartContract struct {
	contractapi.Contract
}

// 注册用户
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

// 农产品上链，传入用户ID、农产品上链信息
func (s *SmartContract) Uplink(ctx contractapi.TransactionContextInterface, userID string, traceability_code string, arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 string) (string, error) {
	// 获取用户类型
	userType, err := s.GetUserType(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("failed to get user type: %v", err)
	}

	// 通过溯源码获取农产品的上链信息
	FruitAsBytes, err := ctx.GetStub().GetState(traceability_code)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	// 将农产品的信息转换为Fruit结构体
	var fruit Fruit
	if FruitAsBytes != nil {
		err = json.Unmarshal(FruitAsBytes, &fruit)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal fruit: %v", err)
		}
	}

	//获取时间戳,修正时区
	txtime, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return "", fmt.Errorf("failed to read TxTimestamp: %v", err)
	}
	timeLocation, _ := time.LoadLocation("Asia/Shanghai") // 选择你所在的时区
	time := time.Unix(txtime.Seconds, 0).In(timeLocation).Format("2006-01-02 15:04:05")

	// 获取交易ID
	txid := ctx.GetStub().GetTxID()
	// 给农产品信息中加上溯源码
	fruit.Traceability_code = traceability_code
	// 不同用户类型的上链的参数不一致
	switch userType {
	// 种植户
	case "种植户":
		// 将传入的农产品上链信息转换为Farmer_input结构体
		fruit.Farmer_input.Fa_fruitName = arg1
		fruit.Farmer_input.Fa_origin = arg2
		fruit.Farmer_input.Fa_plantTime = arg3
		fruit.Farmer_input.Fa_pickingTime = arg4
		fruit.Farmer_input.Fa_farmerName = arg5
		fruit.Farmer_input.Fa_img = arg6
		fruit.Farmer_input.Fa_Txid = txid
		fruit.Farmer_input.Fa_Timestamp = time
	// 工厂
	case "工厂":
		// 将传入的农产品上链信息转换为Factory_input结构体
		fruit.Factory_input.Fac_productName = arg1
		fruit.Factory_input.Fac_productionbatch = arg2
		fruit.Factory_input.Fac_productionTime = arg3
		fruit.Factory_input.Fac_factoryName = arg4
		fruit.Factory_input.Fac_contactNumber = arg5
		fruit.Factory_input.Fac_img = arg6
		fruit.Factory_input.Fac_Txid = txid
		fruit.Factory_input.Fac_Timestamp = time
	// 运输司机
	case "运输司机":
		// 将传入的农产品上链信息转换为Driver_input结构体
		fruit.Driver_input.Dr_name = arg1
		fruit.Driver_input.Dr_age = arg2
		fruit.Driver_input.Dr_phone = arg3
		fruit.Driver_input.Dr_carNumber = arg4
		fruit.Driver_input.Dr_transport = arg5
		fruit.Driver_input.Dr_img = arg6
		fruit.Driver_input.Dr_Txid = txid
		fruit.Driver_input.Dr_Timestamp = time
	// 商店
	case "商店":
		// 将传入的农产品上链信息转换为Shop_input结构体
		fruit.Shop_input.Sh_storeTime = arg1
		fruit.Shop_input.Sh_sellTime = arg2
		fruit.Shop_input.Sh_shopName = arg3
		fruit.Shop_input.Sh_shopAddress = arg4
		fruit.Shop_input.Sh_shopPhone = arg5
		fruit.Shop_input.Sh_img = arg6
		fruit.Shop_input.Sh_Txid = txid
		fruit.Shop_input.Sh_Timestamp = time

	}

	//将农产品的信息转换为json格式
	fruitAsBytes, err := json.Marshal(fruit)
	if err != nil {
		return "", fmt.Errorf("failed to marshal fruit: %v", err)
	}
	//将农产品的信息存入区块链
	err = ctx.GetStub().PutState(traceability_code, fruitAsBytes)
	if err != nil {
		return "", fmt.Errorf("failed to put fruit: %v", err)
	}
	//将农产品存入用户的农产品列表
	err = s.AddFruit(ctx, userID, &fruit)
	if err != nil {
		return "", fmt.Errorf("failed to add fruit to user: %v", err)

	}

	return txid, nil
}

// 添加农产品到用户的农产品列表
func (s *SmartContract) AddFruit(ctx contractapi.TransactionContextInterface, userID string, fruit *Fruit) error {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return err
	}
	// 遍历用户的农产品列表，检查是否已经存在该农产品
	for _, existingFruit := range user.FruitList {
		if existingFruit.Traceability_code == fruit.Traceability_code {
			return fmt.Errorf("the fruit with traceability code %s already exists in user %s's fruit list", fruit.Traceability_code, userID)
		}
	}
	// 如果不存在，则将农产品添加到用户的农产品列表中
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

// 获取用户类型
func (s *SmartContract) GetUserType(ctx contractapi.TransactionContextInterface, userID string) (string, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return "", fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return "", err
	}
	return user.UserType, nil
}

// 获取用户信息
func (s *SmartContract) GetUserInfo(ctx contractapi.TransactionContextInterface, userID string) (*User, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return &User{}, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return &User{}, fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// 获取农产品的上链信息
func (s *SmartContract) GetFruitInfo(ctx contractapi.TransactionContextInterface, traceability_code string) (*Fruit, error) {
	FruitAsBytes, err := ctx.GetStub().GetState(traceability_code)
	if err != nil {
		return &Fruit{}, fmt.Errorf("failed to read from world state: %v", err)
	}

	// 将返回的结果转换为Fruit结构体
	var fruit Fruit
	if FruitAsBytes != nil {
		err = json.Unmarshal(FruitAsBytes, &fruit)
		if err != nil {
			return &Fruit{}, fmt.Errorf("failed to unmarshal fruit: %v", err)
		}
		if fruit.Traceability_code != "" {
			return &fruit, nil
		}
	}
	return &Fruit{}, fmt.Errorf("the fruit %s does not exist", traceability_code)
}

// 获取用户的农产品ID列表
func (s *SmartContract) GetFruitList(ctx contractapi.TransactionContextInterface, userID string) ([]*Fruit, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return nil, fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, err
	}
	return user.FruitList, nil
}

// 获取所有的农产品信息
func (s *SmartContract) GetAllFruitInfo(ctx contractapi.TransactionContextInterface) ([]Fruit, error) {
	fruitListAsBytes, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	defer fruitListAsBytes.Close()
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
		// 过滤非农产品的信息
		if fruit.Traceability_code != "" {
			fruits = append(fruits, fruit)
		}
	}
	return fruits, nil
}

// 获取农产品上链历史
func (s *SmartContract) GetFruitHistory(ctx contractapi.TransactionContextInterface, traceability_code string) ([]HistoryQueryResult, error) {
	log.Printf("GetAssetHistory: ID %v", traceability_code)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(traceability_code)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

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
		// 指定目标时区
		targetLocation, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			return nil, err
		}

		// 将时间戳转换到目标时区
		timestamp = timestamp.In(targetLocation)
		// 格式化时间戳为指定格式的字符串
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
