package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

type auth struct {
	Account  string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
	Device   device `json:"device"` // 设备信息
}

// 手机设备信息
type device struct {
	Width         int    `json:"width"`         // 设备屏幕分辨率宽度
	Height        int    `json:"height"`        // 设备屏幕分辨率高度
	BuildId       int    `json:"buildId"`       // 修订版本号，或者诸如"M4-rc20"的标识
	Broad         string `json:"broad"`         // 设备的主板(?)型号
	Brand         string `json:"brand"`         // 与产品或硬件相关的厂商品牌，如"Xiaomi", "Huawei"等
	Device        string `json:"device"`        // 设备在工业设计中的名称
	Model         string `json:"model"`         // 设备型号
	Product       string `json:"product"`       // 整个产品的名称
	Bootloader    string `json:"bootloader"`    // 设备Bootloader的版本
	Hardware      string `json:"hardware"`      // 设备的硬件名称(来自内核命令行或者/proc)
	Fingerprint   string `json:"fingerprint"`   // 构建(build)的唯一标识码
	Serial        string `json:"serial"`        // 硬件序列号
	SdkInt        string `json:"sdkInt"`        // 安卓系统API版本。例如安卓4.4的sdkInt为19
	Incremental   string `json:"incremental"`   // The internal value used by the underlying source control to represent this build. E.g., a perforce changelist number or a git hash.
	Release       string `json:"release"`       // Android系统版本号。例如"5.0", "7.1.1"
	BaseOS        string `json:"baseOS"`        // The base OS build the product is based on.
	SecurityPatch string `json:"securityPatch"` // 安全补丁程序级别
	Codename      string `json:"codename"`      // 开发代号，例如发行版是"REL"
	Imei          string `json:"imei"`          // 设备的IMEI
	AndroidID     string `json:"androidID"`     // 设备的Android ID。Android ID为一个用16进制字符串表示的64位整数，在设备第一次使用时随机生成，之后不会更改，除非恢复出厂设置。
	MacAddress    string `json:"macAddress"`    // 返回设备的Mac地址。该函数需要在有WLAN连接的情况下才能获取，否则会返回null。
}

// @Summary Get Auth
// @Produce  json
// @Param account query string true "account"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	account := c.PostForm("account")
	password := c.PostForm("password")

	a := auth{Account: account, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	merchantModel := models.Merchant{}
	m, err := merchantModel.GetByAccount(account)
	if err != nil {
		appG.Response200(e.SUCCESS, `Can't find this account`, nil)
		return
	}
	if m.Password != password {
		appG.Response200(e.SUCCESS, `Pwd error`, nil)
		return
	}

	token, err := util.GenerateToken(account, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

// func GetAuthBak(c *gin.Context) {
// 	appG := app.Gin{C: c}
// 	valid := validation.Validation{}

// 	username := c.PostForm("username")
// 	password := c.PostForm("password")

// 	a := auth{Username: username, Password: password}
// 	ok, _ := valid.Valid(&a)

// 	if !ok {
// 		app.MarkErrors(valid.Errors)
// 		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
// 		return
// 	}

// 	authService := auth_service.Auth{Username: username, Password: password}
// 	isExist, err := authService.Check()
// 	if err != nil {
// 		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
// 		return
// 	}

// 	if !isExist {
// 		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
// 		return
// 	}

// 	token, err := util.GenerateToken(username, password)
// 	if err != nil {
// 		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
// 		return
// 	}

// 	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
// 		"token": token,
// 	})
// }
