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

func GetLineHandler(ctx *fasthttp.RequestCtx) {
	line_name := string(ctx.QueryArgs().Peek("line_name"))
	db, err := _var.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

    var lines []models.LineModel
    db.Raw("select line_id,line_name from lines where line_name like '%?%'",string(line_name)).Find(&lines)

	result,err := json.Marshal(lines)
	if err != nil {
		ctx.Write([]byte("emptyreal"))
	}
	ctx.Write([]byte(result))
}