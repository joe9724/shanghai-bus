package handlers

import (
	"github.com/valyala/fasthttp"
	"fmt"
	"github.com/json-iterator/go"
	_"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	_var"shanghai-bus/var"
	"shanghai-bus/models"
)
var jsonffff = jsoniter.ConfigCompatibleWithStandardLibrary
func GetNearHandler(ctx *fasthttp.RequestCtx) {
	lat := string(ctx.QueryArgs().Peek("st_lat"))
	lon := string(ctx.QueryArgs().Peek("st_lon"))
	db, err := _var.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

    var stations []models.StationModel
    sql := "select distinct(st_name) from line_station where "+
	" st_lat > ?-1 and "+
	" st_lat < ?+1 and "+
	" st_lon > ?-1 and "+
	" st_lon < ?+1 "+
	" order by ACOS(SIN((? * 3.1415) / 180 ) *SIN((st_lat * 3.1415) / 180 ) +COS((? * 3.1415) / 180 ) * COS((st_lat * 3.1415) / 180 ) *COS((?* 3.1415) / 180 - (st_lon * 3.1415) / 180 ) ) * 6380 asc limit 10"

	//fmt.Println("sql is,", sql,lat,lon)

	db.Raw(sql,lat,lat,lon,lon,lat,lat,lon).Find(&stations)

	//循环计算每个站台经过的线路
	for i:=0; i<len(stations);i++  {
		var passlines []models.LineModel
		db.Raw("select line_station.line_id,`lines`.line_name from line_station left join `lines` on line_station.line_id=`lines`.line_id where st_name=? GROUP BY line_id,line_name",stations[i].StName).Find(&passlines)
		stations[i].PassLines = passlines
	}

	result,err := jsonffff.Marshal(stations)
	if err != nil {
		ctx.Write([]byte("emptyreal"))
	}
	ctx.Write([]byte(result))
}