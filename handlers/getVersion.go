package handlers

import (
	_"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/valyala/fasthttp"
	"encoding/json"
	"shanghai-bus/models"
	_var "shanghai-bus/var"
	"fmt"
	"strings"
)

func GetVersionHandler(ctx *fasthttp.RequestCtx) {

	client := string(ctx.QueryArgs().Peek("client"))
	client = strings.ToLower(client)
	if len(client) == 0 {
		ctx.Write([]byte(_var.Response200(301, "Client参数未传或传入错误")))
		return
	}

	db, err := _var.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var version models.Version
	sql := "SELECT number, content, download, forced, version_url, about_url, client FROM version WHERE client = ? AND status = 0"
	db.Raw(sql, client).First(&version)

	var splash models.Splash
	sql = "SELECT title, pic, pic_x, web_url, count_down FROM splash WHERE client = ? AND status = 0"
	db.Raw(sql, client).First(&splash)

	var data models.VersionData
	data.Version = version
	data.Splash = splash

	result, err := json.Marshal(data)
	if err != nil {
		ctx.Write([]byte(_var.Response200(302, "当前没有版本号")))
	}
	ctx.Write([]byte(result))
}
