package model

type MysqlUser struct {
    UserID   string `json:"userID"`
    Username string `json:"username"`
    Password string `json:"password"`
    RealInfo string `json:"realInfo"`
}

type User struct {
    UserID       string     `json:"userID"`
    UserType     string     `json:"userType"`
    RealInfoHash string     `json:"realInfoHash"`
    ProductList  []*Product `json:"productList"`
}

type Product struct {
    TraceabilityCode     string               `json:"traceabilityCode"`
    BeeFarmInput         BeeFarmInput         `json:"beeFarmInput"`
    ProcessingPlantInput ProcessingPlantInput `json:"processingPlantInput"`
    WholesalerInput      WholesalerInput      `json:"wholesalerInput"`
    RetailerInput        RetailerInput        `json:"retailerInput"`
}

type HistoryQueryResult struct {
    Record    *Product `json:"record"`
    TxId      string   `json:"txId"`
    Timestamp string   `json:"timestamp"`
    IsDelete  bool     `json:"isDelete"`
}

type BeeFarmInput struct {
    BeeFarmName      string `json:"beeFarmName"`
    BeeFarmLocation  string `json:"beeFarmLocation"`
    BeeBoxId         string `json:"beeBoxId"`
    HoneyVariety     string `json:"honeyVariety"`
    FlowerVariety    string `json:"flowerVariety"`
    BeeFarmTxid      string `json:"beeFarmTxid"`
    BeeFarmTimestamp string `json:"beeFarmTimestamp"`
}

type ProcessingPlantInput struct {
    ProcessingPlantName      string `json:"processingPlantName"`
    ProcessingPlantLocation  string `json:"processingPlantLocation"`
    ProcessingBatchId        string `json:"processingBatchId"`
    PackagingSpecification   string `json:"packagingSpecification"`
    ShelfLife                string `json:"shelfLife"`
    ProcessingPlantTxid      string `json:"processingPlantTxid"`
    ProcessingPlantTimestamp string `json:"processingPlantTimestamp"`
}

type WholesalerInput struct {
    WarehouseName        string `json:"warehouseName"`
    WarehouseLocation    string `json:"warehouseLocation"`
    WholesalerBatchId    string `json:"wholesalerBatchId"`
    TransportationMethod string `json:"transportationMethod"`
    TransportMode        string `json:"transportMode"`
    WholesalerTxid       string `json:"wholesalerTxid"`
    WholesalerTimestamp  string `json:"wholesalerTimestamp"`
}

type RetailerInput struct {
    StoreName         string `json:"storeName"`
    StoreLocation     string `json:"storeLocation"`
    RetailerBatchId   string `json:"retailerBatchId"`
    SalesChannel      string `json:"salesChannel"`
    SalesPrice        string `json:"salesPrice"`
    RetailerTxid      string `json:"retailerTxid"`
    RetailerTimestamp string `json:"retailerTimestamp"`
}
