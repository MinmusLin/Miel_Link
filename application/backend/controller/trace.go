package controller

import (
    "backend/pkg"
    "fmt"

    "github.com/gin-gonic/gin"
)

func Uplink(c *gin.Context) {
    farmerTraceabilityCode := pkg.GenerateID()
    args := BuildArgs(c, farmerTraceabilityCode)
    res, err := pkg.ChaincodeInvoke("Uplink", args)
    if err != nil {
        c.JSON(200, gin.H{
            "message": "Uplink failed: " + err.Error(),
        })
        return
    }
    c.JSON(200, gin.H{
        "code":              200,
        "message":           "Uplink successfully",
        "txid":              res,
        "traceability_code": farmerTraceabilityCode,
    })
}

func GetFruitInfo(c *gin.Context) {
    res, err := pkg.ChaincodeQuery("GetFruitInfo", c.PostForm("traceability_code"))
    if err != nil {
        c.JSON(200, gin.H{
            "message": "Query failed: " + err.Error(),
        })
    }
    c.JSON(200, gin.H{
        "code":    200,
        "message": "Query successfully",
        "data":    res,
    })
}

func GetFruitList(c *gin.Context) {
    userID, _ := c.Get("userID")
    res, err := pkg.ChaincodeQuery("GetFruitList", userID.(string))
    if err != nil {
        c.JSON(200, gin.H{
            "message": "Query failed: " + err.Error(),
        })
    }
    c.JSON(200, gin.H{
        "code":    200,
        "message": "Query successfully",
        "data":    res,
    })
}

func GetAllFruitInfo(c *gin.Context) {
    res, err := pkg.ChaincodeQuery("GetAllFruitInfo", c.PostForm("arg"))
    if err != nil {
        c.JSON(200, gin.H{
            "message": "Query failed: " + err.Error(),
        })
    }
    c.JSON(200, gin.H{
        "code":    200,
        "message": "Query successfully",
        "data":    res,
    })
}

func GetFruitHistory(c *gin.Context) {
    res, err := pkg.ChaincodeQuery("GetFruitHistory", c.PostForm("traceability_code"))
    if err != nil {
        c.JSON(200, gin.H{
            "message": "Query failed: " + err.Error(),
        })
    }
    c.JSON(200, gin.H{
        "code":    200,
        "message": "Query successfully",
        "data":    res,
    })
}

func BuildArgs(c *gin.Context, farmerTraceabilityCode string) []string {
    var args []string
    userID, _ := c.Get("userID")
    userType, _ := pkg.ChaincodeQuery("GetUserType", userID.(string))
    args = append(args, userID.(string))
    fmt.Print(userID)
    if userType == "种植户" {
        args = append(args, farmerTraceabilityCode)
    } else {
        args = append(args, c.PostForm("traceability_code"))
    }
    args = append(args, c.PostForm("arg1"))
    args = append(args, c.PostForm("arg2"))
    args = append(args, c.PostForm("arg3"))
    args = append(args, c.PostForm("arg4"))
    args = append(args, c.PostForm("arg5"))
    args = append(args, c.PostForm("arg6"))
    args = append(args, c.PostForm("arg7"))
    return args
}
