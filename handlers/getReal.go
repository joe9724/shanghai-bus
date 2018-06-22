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
	"strconv"
)
var jsonf = jsoniter.ConfigCompatibleWithStandardLibrary
func GetRealHandler(ctx *fasthttp.RequestCtx) {
	lineID := ctx.QueryArgs().Peek("line_id")
	updown := ctx.QueryArgs().Peek("updown")
	//查询数据库获取updown对应的segmentid
	var segstr int64
	db, err := _var.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	db.Raw("select seg_id from line_seg_relation where line_id=? and updown=?",string(lineID),string(updown)).Count(&segstr)
	fmt.Println("select seg_id from line_seg_relation where line_id=? and updown=?",string(lineID),string(updown))
	fmt.Println("segstr is",segstr)

	resp, err := http.Get("http://61.132.47.90:8998/BusService/Query_ByRouteID/?RouteID="+string(lineID)+"&Segmentid="+strconv.FormatInt(segstr,10) )
	fmt.Println("sql is:","http://61.132.47.90:8998/BusService/Query_ByRouteID/?RouteID="+string(lineID)+"&Segmentid="+strconv.FormatInt(segstr,10) )
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	var data models.Real
	//
	jsonf.Unmarshal(body, &data)
	fmt.Println("data is",data)
    var real []models.BtkReal
    fmt.Println(len(data.RStaRealTInfoList))
	for i:=0;i<len(data.RStaRealTInfoList) ;i++  {
		var temp models.BtkReal
		temp.Name = data.RStaRealTInfoList[i].MediaRouteName
		temp.Type = data.RStaRealTInfoList[i].BusType
		temp.EAnum = data.RStaRealTInfoList[i].ExpArriveBusStaNum
		temp.RsNum = data.RStaRealTInfoList[i].RStanum
		temp.StID = data.RStaRealTInfoList[i].StationID
		temp.StopNum = data.RStaRealTInfoList[i].StopBusStaNum
		fmt.Println("stopNum is:",data.RStaRealTInfoList[i].StopBusStaNum)
        real = append(real,temp)
	}
	fmt.Println(real)
	result,err := json.Marshal(real)
	if err != nil {
		ctx.Write([]byte("emptyreal"))
	}
	ctx.Write([]byte(result))
}