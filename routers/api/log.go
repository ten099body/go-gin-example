package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
)

type ReqReceive struct {
	Key  string   `json:"key"`
	Data []string `json:"data"`
}

var pathClientLog = setting.AppSetting.RuntimeRootPath + setting.AppSetting.PathClientLog

// 接收客户端传递过来的日志
func LogReceive(c *gin.Context) {
	// 读取row格式请求体数据
	b, _ := c.GetRawData()
	// 定义map或结构体
	var m ReqReceive
	// 反序列化
	_ = json.Unmarshal(b, &m)

	// fmt.Println("检查变量类型", reflect.TypeOf(m["data"]))

	filePath := pathClientLog
	fileName := m.Key
	file, err := os.OpenFile(filePath+fileName, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		c.JSON(http.StatusOK, "Save failed")
		return
	}
	//及时关闭
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < len(m.Data); i++ {
		writer.WriteString(m.Data[i] + "\n")
	}
	writer.Flush()
	c.JSON(http.StatusOK, "Success")
}

// 显示日志列表
func LogList(c *gin.Context) {
	pwd := pathClientLog
	fmt.Println("pwd", pwd)
	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	var fileNames []string
	for i := range fileInfoList {
		fileNames = append(fileNames, fileInfoList[i].Name())
	}

	data := map[string]interface{}{
		"fileNames": fileNames,
	}
	c.HTML(http.StatusOK, "log/list.tmpl", data)
}

func LogDetail(c *gin.Context) {
	name := c.Param("name")
	content, err := ioutil.ReadFile(pathClientLog + name)
	if err != nil {
		c.JSON(http.StatusOK, "Error")
		return
	}
	data := map[string]interface{}{
		"title":   name,
		"content": string(content),
	}
	c.HTML(http.StatusOK, "log/detail.html", data)
}
