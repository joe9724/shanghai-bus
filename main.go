package main

import (
	"github.com/valyala/fasthttp"
    "github.com/json-iterator/go"
	"fmt"
	handler "shanghai-bus/handlers"

)
var json = jsoniter.ConfigCompatibleWithStandardLibrary


func main() {

	// the corresponding fasthttp code
	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/getUpdown":
			handler.GetUpDownHandler(ctx)
		case "/getStation":
			handler.GetStationHandler(ctx)
		case "/getReal":
			handler.GetRealHandler(ctx)
		case "/getJoke":
			handler.GetJokeHandler(ctx)
		case "/getNear":
			handler.GetNearHandler(ctx)
		case "/getLine/search":
			handler.GetLineHandler(ctx)
		case "/getLinePassStation":
			handler.GetLinePassStationHandler(ctx)
		case "/getVersion":
			handler.GetVersionHandler(ctx)
		case "/getSearch":
			handler.GetSearchHandler(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}
	fasthttp.ListenAndServe(":523", m)
	fmt.Println("start server...at port 523")
}
