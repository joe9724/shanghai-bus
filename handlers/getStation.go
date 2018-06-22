package handlers

import (
	"github.com/valyala/fasthttp"
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"shanghai-bus/models"
	_var"shanghai-bus/var"
	_"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

func GetStationHandler(ctx *fasthttp.RequestCtx) {
	upDown := ctx.QueryArgs().Peek("updown")
	lineID := ctx.QueryArgs().Peek("line_id")
	resp, err := http.Get("http://61.132.47.90:8998/BusService/Require_RouteStatData/?RouteID=" + string(lineID))
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	//fmt.Println(string(body))
	var data []models.LineStationModel
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(body, &data)
	if len(data)>0{
		temp := data[0]
		var ls models.BtkLineStation
		ls.ID = temp.RouteID
		ls.Name = temp.RouteName
		ls.Dir = temp.RouteType

		db, err := _var.OpenConnection()
		if err != nil {
			fmt.Println(err.Error())
		}
		defer db.Close()

		if len(temp.SegmentList)==0{
			ctx.Write([]byte("emptyDir"))

		}else if len(temp.SegmentList) == 1{
			//只有一个行驶方向
            //判断是否已有line_seg映射关系

            fmt.Println("line_id is",string(lineID))

			var seg0 models.Seg
			seg0.Info =temp.SegmentList[0].FirstLastShiftInfo
			seg0.Price = temp.SegmentList[0].RoutePrice
			seg0.SegID = temp.SegmentList[0].SegmentID
			seg0.SegName = temp.SegmentList[0].SegmentName
			//重写stations
			var tempstation models.BtkStation
			for i:=0; i<len(temp.SegmentList[0].StationList);i++  {
				fmt.Println(temp.SegmentList[0].StationList[i].StationPosition.Longitude)
				tempstation.Lat = temp.SegmentList[0].StationList[i].StationPosition.Latitude
				tempstation.Lon = temp.SegmentList[0].StationList[i].StationPosition.Longitude
				fmt.Println("lat lon is:",tempstation.Lat,tempstation.Lon)
				tempstation.StID = temp.SegmentList[0].StationList[i].StationID
				tempstation.StName = temp.SegmentList[0].StationList[i].StationName
				tempstation.StNo = temp.SegmentList[0].StationList[i].StationNo
				seg0.Stations = append(seg0.Stations,tempstation)
			}
			ls.SEG = append(ls.SEG,seg0)

			var id int64
			if string(upDown) == "0"{
				db.Raw("select id from line_seg_relation where line_id=? and updown=?",string(lineID),string(upDown)).Count(&id)
				fmt.Println("id is",id)
				if id==0 {
					//db.Exec("insert into line_seg_relation(line_id,updown,seg_id,start_station,end_station,summary) values(?,?,?,?,?,?)",string(lineID),string(upDown),ls.SEG[0].SegID,temp.SegmentList[0].StationList[0].StationName,temp.SegmentList[0].StationList[len(temp.SegmentList[0].StationList)-1].StationName,temp.SegmentList[0].FirstTime+"至"+temp.SegmentList[0].LastTime)
					fmt.Println("insert into line_seg_relation(line_id,updown,seg_id,start_station,end_station) values(?,?,?,?,?)",string(lineID),string(upDown),ls.SEG[0].SegID,temp.SegmentList[0].StationList[0].StationName,temp.SegmentList[0].StationList[len(temp.SegmentList[0].StationList)-1].StationName)
				}
				//插入线路站台表
				//for sn:=0; sn<len(temp.SegmentList[0].StationList);sn++  {
					//db.Exec("insert into line_station(line_id,line_updown,st_name,st_lat,st_lon) values(?,?,?,?,?)",string(lineID),string(upDown),temp.SegmentList[0].StationList[sn].StationName,temp.SegmentList[0].StationList[sn].StationPosition.Latitude,temp.SegmentList[0].StationList[sn].StationPosition.Longitude)
				//}
			}

		}else if len(temp.SegmentList) == 2{
			var seg0 models.Seg
			seg0.Info =temp.SegmentList[0].FirstLastShiftInfo
			seg0.Price = temp.SegmentList[0].RoutePrice
			seg0.SegID = temp.SegmentList[0].SegmentID
			seg0.SegName = temp.SegmentList[0].SegmentName
			//重写stations
			var tempstation models.BtkStation
			for i:=0; i<len(temp.SegmentList[0].StationList);i++  {
				fmt.Println(temp.SegmentList[0].StationList[i].StationPosition.Longitude)
				tempstation.Lat = temp.SegmentList[0].StationList[i].StationPosition.Latitude
				tempstation.Lon = temp.SegmentList[0].StationList[i].StationPosition.Longitude
				tempstation.StID = temp.SegmentList[0].StationList[i].StationID
				tempstation.StName = temp.SegmentList[0].StationList[i].StationName
				tempstation.StNo = temp.SegmentList[0].StationList[i].StationNo
				seg0.Stations = append(seg0.Stations,tempstation)
				fmt.Println("lat lon is:",tempstation.Lat,tempstation.Lon)
			}
			ls.SEG = append(ls.SEG,seg0)

			var id int64
			if string(upDown) == "0"{
				db.Raw("select id from line_seg_relation where line_id=? and updown=?",string(lineID),string(upDown)).Count(&id)
				fmt.Println("id is",id)
				if id==0 {
					//db.Exec("insert into line_seg_relation(line_id,updown,seg_id,start_station,end_station,summary) values(?,?,?,?,?,?)",string(lineID),string(upDown),ls.SEG[0].SegID,temp.SegmentList[0].StationList[0].StationName,temp.SegmentList[0].StationList[len(temp.SegmentList[0].StationList)-1].StationName,temp.SegmentList[0].FirstTime+"至"+temp.SegmentList[0].LastTime)
					fmt.Println("insert into line_seg_relation(line_id,updown,seg_id,start_station,end_station) values(?,?,?,?,?)",string(lineID),string(upDown),ls.SEG[0].SegID,temp.SegmentList[0].StationList[0].StationName,temp.SegmentList[0].StationList[len(temp.SegmentList[0].StationList)-1].StationName)
				}
				//for sn:=0; sn<len(temp.SegmentList[0].StationList);sn++  {
					//db.Exec("insert into line_station(line_id,line_updown,st_name,st_lat,st_lon) values(?,?,?,?,?)",string(lineID),string(upDown),temp.SegmentList[0].StationList[sn].StationName,temp.SegmentList[0].StationList[sn].StationPosition.Latitude,temp.SegmentList[0].StationList[sn].StationPosition.Longitude)
				//}


			}


			//
			var seg1 models.Seg
			seg1.Info =temp.SegmentList[1].FirstLastShiftInfo
			seg1.Price = temp.SegmentList[1].RoutePrice
			seg1.SegID = temp.SegmentList[1].SegmentID
			seg1.SegName = temp.SegmentList[1].SegmentName
			//重写stations
			var tempstation1 models.BtkStation
			for i:=1; i<len(temp.SegmentList[1].StationList);i++  {
				tempstation1.Lat = temp.SegmentList[1].StationList[i].StationPosition.Latitude
				tempstation1.Lon = temp.SegmentList[1].StationList[i].StationPosition.Longitude
				tempstation1.StID = temp.SegmentList[1].StationList[i].StationID
				tempstation1.StName = temp.SegmentList[1].StationList[i].StationName
				tempstation1.StNo = temp.SegmentList[1].StationList[i].StationNo
				seg1.Stations = append(seg1.Stations,tempstation1)
				fmt.Println("lat lon is:",tempstation.Lat,tempstation.Lon)
			}
			ls.SEG = append(ls.SEG,seg1)

			//
			var id2 int64
			if string(upDown) == "1"{
				db.Raw("select id from line_seg_relation where line_id=? and updown=?",string(lineID),string(upDown)).Count(&id2)
				if id2==0 {
					//db.Exec("insert into line_seg_relation(line_id,updown,seg_id,start_station,end_station,summary) values(?,?,?,?,?,?)",string(lineID),string(upDown),ls.SEG[1].SegID,temp.SegmentList[1].StationList[0].StationName,temp.SegmentList[1].StationList[len(temp.SegmentList[1].StationList)-1].StationName,temp.SegmentList[1].FirstTime+"至"+temp.SegmentList[1].LastTime)

				}
				//for sn:=0; sn<len(temp.SegmentList[1].StationList);sn++  {
					//db.Exec("insert into line_station(line_id,line_updown,st_name,st_lat,st_lon) values(?,?,?,?,?)",string(lineID),string(upDown),temp.SegmentList[1].StationList[sn].StationName,temp.SegmentList[1].StationList[sn].StationPosition.Latitude,temp.SegmentList[1].StationList[sn].StationPosition.Longitude)
				//}
			}

		}
		result,err := json.Marshal(ls)
		if err != nil {
			ctx.Write([]byte("emptyres"))
		}
		ctx.Write([]byte(result))

	}else{
		ctx.Write([]byte("empty"))
	}
}