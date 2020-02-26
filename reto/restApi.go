package main

import (
	"fmt"
	"log"

	"model"

	"github.com/buaazp/fasthttprouter"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

var theModel model.LogicModel

var item model.Item
var domain model.Domain

func main() {
	router := fasthttprouter.New()

	//Endpoints
	router.GET("/infoServers/:domain", getServersEndpoint)
	router.GET("/listedServers", getListedServers)
	log.Fatal(fasthttp.ListenAndServe(":3000", CORS(router.Handler)))

}

func getServersEndpoint(request *fasthttp.RequestCtx) {
	params := request.UserValue("domain")
	domainToSearch := fmt.Sprint(params)

	domain = theModel.GetInformationFromServers(domainToSearch)

	domain = theModel.GetInformationFromDomain(domain, domainToSearch)

	var NewJson = jsoniter.Config{
		TagKey: "newtag",
	}.Froze()

	jsonToPrint := NewJson.NewEncoder(request)
	jsonToPrint.SetIndent("", "\t")
	jsonToPrint.Encode(domain)
}

func getListedServers(request *fasthttp.RequestCtx) {
	item := theModel.GetListedServers()
	var NewJson = jsoniter.Config{
		TagKey: "newtag",
	}.Froze()

	jsonToPrint := NewJson.NewEncoder(request)
	jsonToPrint.SetIndent("", "\t")
	jsonToPrint.Encode(item)
}

var (
	corsAllowHeaders     = "authorization"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)

		next(ctx)
	}
}
