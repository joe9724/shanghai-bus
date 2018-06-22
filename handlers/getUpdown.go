package handlers

import (
	"github.com/valyala/fasthttp"
	"fmt"
	"github.com/json-iterator/go"
	"encoding/json"
	_"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	_var"shanghai-bus/var"
	"shanghai-bus/models"
)
var jsonfff = jsoniter.ConfigCompatibleWithStandardLibrary
func GetUpDownHandler(ctx *fasthttp.RequestCtx) {
	line_id := string(ctx.QueryArgs().Peek("line_id"))
	db, err := _var.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

    var updowns []models.UpDown
    db.Raw("select updown as up_down,start_station,end_station from line_seg_relation where line_id=?",line_id).Find(&updowns)

	result,err := json.Marshal(updowns)
	if err != nil {
		ctx.Write([]byte("emptyreal"))
	}
	ctx.Write([]byte(result))
}