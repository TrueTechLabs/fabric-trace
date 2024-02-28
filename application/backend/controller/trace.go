package controller

import (
	"backend/pkg"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 保存了农产品上链与查询的函数

func Uplink(c *gin.Context) {
	farmer_traceability_code := pkg.GenerateID()
	args := buildArgs(c, farmer_traceability_code)
	res, err := pkg.ChaincodeInvoke("Uplink", args)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "uplink failed" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":              200,
		"message":           "uplink success",
		"txid":              res,
		"traceability_code": farmer_traceability_code,
	})
}

// 获取农产品的上链信息
func GetFruitInfo(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetFruitInfo", c.PostForm("traceability_code"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed" + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

// 获取用户的农产品ID列表
func GetFruitList(c *gin.Context) {
	userID, _ := c.Get("userID")
	res, err := pkg.ChaincodeQuery("GetFruitList", userID.(string))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed" + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

// 获取所有的农产品信息
func GetAllFruitInfo(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetAllFruitInfo", c.PostForm("arg"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed" + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

// 获取农产品上链历史
// func (s *SmartContract) GetFruitHistory(ctx contractapi.TransactionContextInterface, traceability_code string) ([]HistoryQueryResult, error) {
func GetFruitHistory(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetFruitHistory", c.PostForm("traceability_code"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed" + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

func buildArgs(c *gin.Context, farmer_traceability_code string) []string {
	var args []string
	userID, _ := c.Get("userID")
	userType, _ := pkg.ChaincodeQuery("GetUserType", userID.(string))
	args = append(args, userID.(string))
	fmt.Print(userID)
	// 种植户不需要输入溯源码，其他用户需要，通过雪花算法生成ID
	if userType == "种植户" {
		args = append(args, farmer_traceability_code)
	} else {
		args = append(args, c.PostForm("traceability_code"))
	}
	args = append(args, c.PostForm("arg1"))
	args = append(args, c.PostForm("arg2"))
	args = append(args, c.PostForm("arg3"))
	args = append(args, c.PostForm("arg4"))
	args = append(args, c.PostForm("arg5"))
	return args
}
