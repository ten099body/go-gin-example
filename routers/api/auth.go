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
