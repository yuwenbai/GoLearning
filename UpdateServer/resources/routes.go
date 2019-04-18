package main

import (
	"net/http"
)

//Route 路由实例
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes 路由实例数组
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"handerGetFile",
		"GET",
		"/downloadFile",
		handerGetFile,
	},
	Route{
		"GetCurrentVersion",
		"POST",
		"/GetVersion",
		GetCurrentVersion,
	},
	Route{
		"GetAllPackageNames",
		"GET",
		"/GetAllPackageNames",
		GetAllPackageNames,
	},
	Route{
		"GeneratePatch",
		"GET",
		"/GeneratePatch",
		GeneratePatch,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"CheckUpdateInfoJson",
		"GET",
		"/CheckUpdateInfoJson",
		CheckUpdateInfoJSON,
	},
	Route{
		"uploadIOSHandler",
		"POST",
		"/uploadIOS",
		uploadIOSHandler,
	},
	Route{
		"uploadAndroidHandler",
		"POST",
		"/uploadAndroid",
		uploadAndroidHandler,
	},
}
