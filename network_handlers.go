package main

import (	
	"github.com/valyala/fasthttp"
	
	"SysSpecter/network"
	
	"fmt"
	"os"
)


var handel_get_connections = func(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)

	result := network.GetEstablishedConnections()
	json_response, err := Json(result)

	ctx_error_handler(ctx, err)
	fmt.Fprintf(ctx, json_response)
}

var handel_get_per_second = func(ctx *fasthttp.RequestCtx) {
	result := map[string]interface{}{
		"incoming": network.IncomingPerSecond,
		"outgoing": network.OutgoingPerSecond,
	}
	json_response, err := Json(result)
	ctx_error_handler(ctx, err)
	fmt.Fprintf(ctx, json_response)
}

var handel_get_interface_list = func(ctx *fasthttp.RequestCtx) {
	json_response, err := Json(network.ListInterfaces())
	ctx_error_handler(ctx, err)
	fmt.Fprintf(ctx, json_response)
}

// RX Part handlers
var handel_get_income_bytes = func(ctx *fasthttp.RequestCtx) {
	query_args := ctx.QueryArgs()
	if !query_args.Has("interface") {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "Missing 'interface' query parameter"}`)
		return
	}

	interface_chosen := string(query_args.Peek("interface"))
	if !Contains(network.ListInterfaces(), interface_chosen) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "'interface' incorrect"}`)
		return
	}
	result := map[string]interface{}{
		"bytes": network.RXBytes(interface_chosen),
	}
	json_response, err := Json(result)
	ctx_error_handler(ctx, err)
	fmt.Fprintf(ctx, json_response)
}

var handel_get_income_packets = func(ctx *fasthttp.RequestCtx) {
	query_args := ctx.QueryArgs()
	if !query_args.Has("interface") {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "Missing 'interface' query parameter"}`)
		return
	}

	interface_chosen := string(query_args.Peek("interface"))
	if !Contains(network.ListInterfaces(), interface_chosen) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "'interface' incorrect"}`)
		return
	}
	result := map[string]interface{}{
		"count": network.RXPackets(interface_chosen),
	}
	json_response, err := Json(result)
	ctx_error_handler(ctx, err)
	fmt.Fprintf(ctx, json_response)
}

var handel_get_income_dropped = func(ctx *fasthttp.RequestCtx) {
	query_args := ctx.QueryArgs()
	if !query_args.Has("interface") {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "Missing 'interface' query parameter"}`)
		return
	}

	interface_chosen := string(query_args.Peek("interface"))
	if !Contains(network.ListInterfaces(), interface_chosen) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "'interface' incorrect"}`)
		return
	}
	result := map[string]interface{}{
		"count": network.RXDropped(interface_chosen),
	}
	json_response, err := Json(result)
	ctx_error_handler(ctx, err)
	fmt.Fprintf(ctx, json_response)
}

var handel_get_income_erros = func(ctx *fasthttp.RequestCtx) {
	query_args := ctx.QueryArgs()
	if !query_args.Has("interface") {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "Missing 'interface' query parameter"}`)
		return
	}

	interface_chosen := string(query_args.Peek("interface"))
	if !Contains(network.ListInterfaces(), interface_chosen) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "'interface' incorrect"}`)
		return
	}
	result := map[string]interface{}{
		"count": network.RXErrors(interface_chosen),
	}
	json_response, err := Json(result)
	ctx_error_handler(ctx, err)
	fmt.Fprintf(ctx, json_response)
}

// TX Part Handlers

var handel_get_outgo_bytes = func(ctx *fasthttp.RequestCtx) {
	query_args := ctx.QueryArgs()
	if !query_args.Has("interface") {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "Missing 'interface' query parameter"}`)
		return
	}

	interface_chosen := string(query_args.Peek("interface"))
	if !Contains(network.ListInterfaces(), interface_chosen) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "'interface' incorrect"}`)
		return
	}

	result := map[string]interface{}{
		"bytes": network.TXBytes(interface_chosen),
	}
	json_response, err := Json(result)
	ctx_error_handler(ctx, err)
	fmt.Fprintf(ctx, json_response)
}

var handel_get_outgo_packets = func(ctx *fasthttp.RequestCtx) {
	query_args := ctx.QueryArgs()
	if !query_args.Has("interface") {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "Missing 'interface' query parameter"}`)
		return
	}

	interface_chosen := string(query_args.Peek("interface"))
	if !Contains(network.ListInterfaces(), interface_chosen) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "'interface' incorrect"}`)
		return
	}
	result := map[string]interface{}{
		"count": network.TXPackets(interface_chosen),
	}
	json_response, err := Json(result)
	ctx_error_handler(ctx, err)
	fmt.Fprintf(ctx, json_response)
}

var handel_get_outgo_dropped = func(ctx *fasthttp.RequestCtx) {
	query_args := ctx.QueryArgs()
	if !query_args.Has("interface") {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "Missing 'interface' query parameter"}`)
		return
	}

	interface_chosen := string(query_args.Peek("interface"))
	if !Contains(network.ListInterfaces(), interface_chosen) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "'interface' incorrect"}`)
		return
	}
	result := map[string]interface{}{
		"count": network.TXDropped(interface_chosen),
	}
	json_response, err := Json(result)
	ctx_error_handler(ctx, err)
	fmt.Fprintf(ctx, json_response)
}

var handel_get_outgo_erros = func(ctx *fasthttp.RequestCtx) {
	query_args := ctx.QueryArgs()
	if !query_args.Has("interface") {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "Missing 'interface' query parameter"}`)
		return
	}

	interface_chosen := string(query_args.Peek("interface"))
	if !Contains(network.ListInterfaces(), interface_chosen) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, `{"error": "'interface' incorrect"}`)
		return
	}
	result := map[string]interface{}{
		"count": network.TXErrors(interface_chosen),
	}
	json_response, err := Json(result)
	ctx_error_handler(ctx, err)
	fmt.Fprintf(ctx, json_response)
}

var handel_get_graph = func(ctx *fasthttp.RequestCtx) {
	pngData, err := network.NetworkData_start.RenderChartPNG()
	if err != nil {
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.Header.Set("Content-Type", "image/png")
	ctx.Write(pngData)
}


var handel_live_graph = func (ctx *fasthttp.RequestCtx)  {
	ctx.SetContentType("text/html; charset=utf-8")
	data,_ := os.ReadFile("templates/index.html")
	fmt.Fprintf(ctx, string(data))
}