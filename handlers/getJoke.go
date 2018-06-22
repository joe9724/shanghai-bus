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
	"strconv"
)
var jsonff = jsoniter.ConfigCompatibleWithStandardLibrary
func GetJokeHandler(ctx *fasthttp.RequestCtx) {
	pageSize := string(ctx.QueryArgs().Peek("pageSize")[:])
	fmt.Println("pageSize is",pageSize)
	pageIndex := string(ctx.QueryArgs().Peek("pageIndex")[:])
	db, err := _var.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var datalist []models.Joke
	sql := "select id,content,DATE_FORMAT(create_time,'%Y-%m-%d %T')as create_time from joke order by id desc"
	size, _ := strconv.ParseInt(pageSize, 10, 64)
	index, _ := strconv.ParseInt(pageIndex, 10, 64)

	db.Raw(sql).Limit(size).Offset((index-1)*size).Find(&datalist)

	result,err := json.Marshal(datalist)
	if err != nil {
		ctx.Write([]byte("emptyreal"))
	}
	ctx.Write([]byte(result))
}