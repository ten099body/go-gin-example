package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
)

type ReqReceive struct {
	Key  string   `json:"key"`
	Data []string `json:"data"`
}

// 接收客户端传递过来的日志
func LogReceive(c *gin.Context) {
	appG := app.Gin{C: c}
	b, _ := c.GetRawData()
	var m ReqReceive
	_ = json.Unmarshal(b, &m)

	// fmt.Println("检查变量类型", reflect.TypeOf(m["data"]))

	fileName := m.Key
	fmt.Println(setting.AppSetting.PathClientLog + fileName)
	file, err := os.OpenFile(setting.AppSetting.PathClientLog+fileName, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		logging.Error(err)
		appG.Response200(e.ERROR, "Save failed", nil)
		return
	}
	//及时关闭
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < len(m.Data); i++ {
		writer.WriteString(m.Data[i] + "\n")
	}
	writer.Flush()
	appG.Response200(e.SUCCESS, "Success", nil)
}

// 显示日志列表
func LogList(c *gin.Context) {
	pwd := setting.AppSetting.PathClientLog
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
	c.HTML(http.StatusOK, "log/list.html", data)
}

func LogDetail(c *gin.Context) {
	name := c.Param("name")
	content, err := ioutil.ReadFile(setting.AppSetting.PathClientLog + name)
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
