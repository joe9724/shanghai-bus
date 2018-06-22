package handlers

import (
	"github.com/valyala/fasthttp"
	"fmt"
	//"github.com/json-iterator/go"
	_"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	_var"shanghai-bus/var"
	"shanghai-bus/models"
	"encoding/json"
)

func GetLinePassStationHandler(ctx *fasthttp.RequestCtx) {
	st_name := string(ctx.QueryArgs().Peek("st_name"))
	db, err := _var.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

    var lines []models.LineModel
    db.Raw("select line_station.line_id,`lines`.line_name from line_station left join `lines` on line_station.line_id=`lines`.line_id where st_name=? GROUP BY line_id,line_name",string(st_name)).Find(&lines)

	result,err := json.Marshal(lines)
	if err != nil {
		ctx.Write([]byte("emptyreal"))
	}
	ctx.Write([]byte(result))
}