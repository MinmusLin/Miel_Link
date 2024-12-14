package chaincode

type User struct {
	UserID       string   `json:"userID"`
	UserType     string   `json:"userType"`
	RealInfoHash string   `json:"realInfoHash"`
	FruitList    []*Fruit `json:"fruitList"`
}

type Fruit struct {
	Traceability_code string        `json:"traceability_code"`
	Farmer_input      Farmer_input  `json:"farmer_input"`
	Factory_input     Factory_input `json:"factory_input"`
	Driver_input      Driver_input  `json:"driver_input"`
	Shop_input        Shop_input    `json:"shop_input"`
}

type HistoryQueryResult struct {
	Record    *Fruit `json:"record"`
	TxId      string `json:"txId"`
	Timestamp string `json:"timestamp"`
	IsDelete  bool   `json:"isDelete"`
}

type Farmer_input struct {
	Fa_fruitName    string `json:"fa_fruitName"`
	Fa_origin       string `json:"fa_origin"`
	Fa_plantTime    string `json:"fa_plantTime"`
	Fa_pickingTime  string `json:"fa_pickingTime"`
	Fa_farmerName   string `json:"fa_farmerName"`
	Fa_Txid         string `json:"fa_txid"`
	Fa_Timestamp    string `json:"fa_timestamp"`
	Fa_IPFSCID      string `json:"fa_ipfscid"`
	Fa_IPFSFileName string `json:"fa_ipfsfilename"`
}

type Factory_input struct {
	Fac_productName     string `json:"fac_productName"`
	Fac_productionbatch string `json:"fac_productionbatch"`
	Fac_productionTime  string `json:"fac_productionTime"`
	Fac_factoryName     string `json:"fac_factoryName"`
	Fac_contactNumber   string `json:"fac_contactNumber"`
	Fac_Txid            string `json:"fac_txid"`
	Fac_Timestamp       string `json:"fac_timestamp"`
	Fac_IPFSCID         string `json:"fac_ipfscid"`
	Fac_IPFSFileName    string `json:"fac_ipfsfilename"`
}

type Driver_input struct {
	Dr_name         string `json:"dr_name"`
	Dr_age          string `json:"dr_age"`
	Dr_phone        string `json:"dr_phone"`
	Dr_carNumber    string `json:"dr_carNumber"`
	Dr_transport    string `json:"dr_transport"`
	Dr_Txid         string `json:"dr_txid"`
	Dr_Timestamp    string `json:"dr_timestamp"`
	Dr_IPFSCID      string `json:"dr_ipfscid"`
	Dr_IPFSFileName string `json:"dr_ipfsfilename"`
}

type Shop_input struct {
	Sh_storeTime    string `json:"sh_storeTime"`
	Sh_sellTime     string `json:"sh_sellTime"`
	Sh_shopName     string `json:"sh_shopName"`
	Sh_shopAddress  string `json:"sh_shopAddress"`
	Sh_shopPhone    string `json:"sh_shopPhone"`
	Sh_Txid         string `json:"sh_txid"`
	Sh_Timestamp    string `json:"sh_timestamp"`
	Sh_IPFSCID      string `json:"sh_ipfscid"`
	Sh_IPFSFileName string `json:"sh_ipfsfilename"`
}
