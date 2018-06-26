package handlers

import (
	"github.com/valyala/fasthttp"
	"net/http"
	"io/ioutil"
	"shanghai-bus/models"
	"fmt"
	"github.com/json-iterator/go"
	"encoding/json"
	_var"shanghai-bus/var"
)
var jsonf = jsoniter.ConfigCompatibleWithStandardLibrary
func GetRealHandler(ctx *fasthttp.RequestCtx) {
	lineID := ctx.QueryArgs().Peek("line_id")
	name := ctx.QueryArgs().Peek("name")
	updown := ctx.QueryArgs().Peek("updown")
	stopID := ctx.QueryArgs().Peek("stopid")

	db, err := _var.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	resp, err := http.Get("http://xxbs.sh.gov.cn:8080/weixinpage/HandlerBus.ashx?action=Three&direction="+string(updown)+"&lineid="+string(lineID)+"&name="+string(name)+"&stopid="+string(stopID))
	fmt.Println("sql is:","http://xxbs.sh.gov.cn:8080/weixinpage/HandlerBus.ashx?action=Three&direction="+string(updown)+"&lineid="+string(lineID)+"&name="+string(name)+"&stopid="+string(stopID))
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	var data models.BtkCar
	//
	jsonf.Unmarshal(body, &data)

	fmt.Println("data is:",data)

	result,err := json.Marshal(data)
	if err != nil {
		ctx.Write([]byte("emptyreal"))
	}
	ctx.Write([]byte(result))
}