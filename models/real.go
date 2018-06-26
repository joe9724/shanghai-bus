package models

type Real struct {
	RouteID int64 `json:"routeId"`
	RunBusNum int64 `json:"runBusNum"`
	RStaRealTInfoList []RealInfo `json:"rStaRealTInfoList"`
	IsEnd string `json:"isEnd"`
}

type RealInfo struct{
	StationID string `json:"stationID"`
	RStanum int64 	`json:"rStanum"`
	ExpArriveBusStaNum int64 `json:"expArriveBusStaNum"`
	StopBusStaNum int64 `json:"stopBusStaNum"`
	BusType int64 `json:"busType"`
	MediaRouteName string `json:"mediaRouteName"`
	islovebus string `json:"islovebus"`
}

type BtkReal struct{
	StID string `json:"sid"`
	RsNum int64 `json:"rsNum"`
	EAnum int64 `json:"eaNum"`
	StopNum int64 `json:"stopNum"`
	Type int64 `json:"type"`
	Name string `json:"name"`

}

type BtkCar struct{
	Cars []Car `json:"cars"`
}

type Car struct{
	Terminal string `json:"terminal"`
	StopDis string `json:"stopdis"`
	Distance string `json:"distance"`
	Time string `json:"time"`

}
