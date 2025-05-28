package controller

import (
	"backend/pkg"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// 保存了农产品上链与查询的函数

func Uplink(c *gin.Context) {
	// 与userID不一样，取ID从第二位作为溯源码，即18位，雪花ID的结构如下（转化为10进制共19位）：
	// +--------------------------------------------------------------------------+
	// | 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
	// +--------------------------------------------------------------------------+
	farmer_traceability_code := pkg.GenerateID()[1:]
	args := buildArgs(c, farmer_traceability_code)
	if args == nil {
		return
	}
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
		"traceability_code": args[1],
	})
}

// 获取农产品的上链信息
func GetFruitInfo(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetFruitInfo", c.PostForm("traceability_code"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "查询失败：" + err.Error(),
		})
		return
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
	res, err := pkg.ChaincodeQuery("GetAllFruitInfo", "")
	fmt.Print("res", res)
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
		// 检查溯源码是否正确
		res, err := pkg.ChaincodeQuery("GetFruitInfo", c.PostForm("traceability_code"))
		if res == "" || err != nil || len(c.PostForm("traceability_code")) != 18 {
			c.JSON(200, gin.H{
				"message": "请检查溯源码是否正确!!",
			})
			return nil
		} else {
			args = append(args, c.PostForm("traceability_code"))
		}
	}
	args = append(args, c.PostForm("arg1"))
	args = append(args, c.PostForm("arg2"))
	args = append(args, c.PostForm("arg3"))
	args = append(args, c.PostForm("arg4"))
	args = append(args, c.PostForm("arg5"))
	file, _ := c.FormFile("file")
	if file != nil {
		// 保存前端传的图片
		c.SaveUploadedFile(file, "files/uploadfiles/"+file.Filename)
		// 计算文件的SHA256哈希值
		fileSHA256, _ := pkg.CalculateFileSHA256("files/uploadfiles/" + file.Filename)
		ext := pkg.GetFileExt(file.Filename)
		c.SaveUploadedFile(file, fmt.Sprintf("files/images/%s.%s", fileSHA256, ext))
		// 清理临时传的文件
		os.Remove("files/uploadfiles/" + file.Filename)
		args = append(args, fmt.Sprintf("%s.%s", fileSHA256, ext))
	}
	args = append(args, "")
	return args
}

func GetImg(c *gin.Context) {
	filename := c.Param("filename")
	filePath := fmt.Sprintf("files/images/%s", filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(404, gin.H{
			"message": "file not found",
		})
		return
	}
	c.File(filePath)
}
