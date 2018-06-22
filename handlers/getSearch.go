package handlers

import (
	_"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/valyala/fasthttp"
	_var "shanghai-bus/var"
	"shanghai-bus/models"
	"fmt"
	"encoding/json"
)

func GetSearchHandler(ctx *fasthttp.RequestCtx) {

	keyword := string(ctx.QueryArgs().Peek("key"))
	if len(keyword) == 0 {
		ctx.Write([]byte(_var.Response200(301, "关键字为必传")))
		return
	}

	db, err := _var.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var lines []models.LineModel
	sql := "SELECT line_id, line_name FROM `lines` WHERE line_name like " + "'" + keyword + "%" + "'" + " order by line_name+1"
	fmt.Println("line sql = ", sql)
	db.Raw(sql).Find(&lines)

	var stations []models.StationModel
	sql = "SELECT st_name FROM line_station WHERE st_name like " + "'" + keyword + "%" + "'" + " group by st_name"
	fmt.Println("station sql = ", sql)
	db.Raw(sql).Find(&stations)

	var data models.SearchRet
	data.Lines = lines
	data.Stations = stations

	result, err := json.Marshal(data)
	if err != nil {
		ctx.Write([]byte(_var.Response200(302, "当前没有记录")))
	}
	ctx.Write([]byte(result))
}
