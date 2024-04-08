package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"net"
	"strconv"
)

var DefaultHeaderPassword = ""
var DefaultIpRestrict = ""

// Handlers part

func handle_limits(ctx *fasthttp.RequestCtx) {
	if DefaultHeaderPassword != "" {
		if string(ctx.Request.Header.Peek("X-Pass")) != DefaultHeaderPassword {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			ctx.SetBodyString("Unauthorized: Incorrect password")
			return
		}
	}

	if DefaultIpRestrict != "" {
		if !ctx.RemoteIP().Equal(net.ParseIP(DefaultIpRestrict)) {
			ctx.SetStatusCode(fasthttp.StatusForbidden)
			ctx.SetBodyString("Forbidden: Access is restricted")
			return
		}
	}

	// Continue with the rest of your handler if the checks pass.
}

var request_handler = func(ctx *fasthttp.RequestCtx) {
	handle_limits(ctx)
	if ctx.Response.StatusCode() != fasthttp.StatusOK {
		// The handle_limits function has already set the response, so we just return.
		return
	}

	ctx.SetContentType("application/json")
	switch string(ctx.Path()) {

	case "/system/status":
		handel_get_system_status(ctx)
	case "/system/memory":
		handel_get_system_mem(ctx)
	case "/system/disks":
		handel_get_system_disk(ctx)

	case "/network/traffic/per_second":
		handel_get_per_second(ctx)
	case "/network/traffic/graph":
		handel_get_graph(ctx)
	case "/network/traffic/live_graph":
		handel_live_graph(ctx)
	case "/network/interface_list":
		handel_get_interface_list(ctx)
	case "/network/connections":
		handel_get_connections(ctx)

	case "/network/incoming/bytes":
		handel_get_income_bytes(ctx)
	case "/network/incoming/packets":
		handel_get_income_packets(ctx)
	case "/network/incoming/errors":
		handel_get_income_erros(ctx)
	case "/network/incoming/dropped":
		handel_get_income_dropped(ctx)

	case "/network/outgoing/bytes":
		handel_get_outgo_bytes(ctx)
	case "/network/outgoing/packets":
		handel_get_outgo_packets(ctx)
	case "/network/outgoing/errors":
		handel_get_outgo_erros(ctx)
	case "/network/outgoing/dropped":
		handel_get_outgo_dropped(ctx)

	default:
		ctx.Error("nothing here, *- check https://github.com/aDarkDev", fasthttp.StatusNotFound)
	}

	statusCode := ctx.Response.StatusCode()
	status_res := ""
	if statusCode == 200{
		status_res = "\033[92m200\033[0m"
	}else{
		status_res = "\033[91m"+strconv.Itoa(statusCode)+"\033[0m"
	}

	ctx.Response.Header.SetServer("SysSpecter/aDarkDev")
	fmt.Println("[+] " + ctx.RemoteIP().To4().String() + " \033[96m" + string(ctx.Method()) + "\033[0m " + string(ctx.URI().RequestURI()) + " - " + status_res)
}

func Run(host string, port int) {
	fasthttp.ListenAndServe(host+":"+strconv.Itoa(port), request_handler)
}
