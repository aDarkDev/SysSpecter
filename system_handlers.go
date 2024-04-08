package main

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/valyala/fasthttp"

	"runtime"
	"time"
	"fmt"

)



var handel_get_system_status = func(ctx *fasthttp.RequestCtx) {
	cpuUsage, _ := cpu.Percent(time.Second, false) // Error ignored for simplicity
	cpuCounts, _ := cpu.Counts(true)               // Error ignored for simplicity

	var memory runtime.MemStats
	runtime.ReadMemStats(&memory)
	virtualMem, _ := mem.VirtualMemory() // Error ignored for simplicity

	uptimeStr,uptimeInt := GetUptime()
	result := map[string]interface{}{
		"cpu_usage": cpuUsage[0],
		"cpu_count": cpuCounts,
		"mem_usage": virtualMem.UsedPercent,
		"mem_size":  float64(virtualMem.Total) / float64(1024*1024),
		"uptime_str": uptimeStr,
		"uptime_int": uptimeInt,
	}

	json_response, err := Json(result)
	ctx_error_handler(ctx, err)
	fmt.Fprintf(ctx, json_response)
}

var handel_get_system_mem = func(ctx *fasthttp.RequestCtx) {
	var memory runtime.MemStats
	runtime.ReadMemStats(&memory)
	virtualMem, _ := mem.VirtualMemory() // Error ignored for simplicity

	result := map[string]interface{}{
		"mem_total_alloc": memory.TotalAlloc,
		"mem_heap_alloc":  memory.HeapAlloc,
		"mem_alloc":       memory.Alloc,
		"mem_swap":        virtualMem.SwapTotal,
		"mem_sys":         memory.Sys,
		"mem_size":        float64(virtualMem.Total) / float64(1024*1024),
	}

	json_response, err := Json(result)
	ctx_error_handler(ctx, err)
	fmt.Fprintf(ctx, json_response)
}

var handel_get_system_disk = func(ctx *fasthttp.RequestCtx) {
	result, err := GetDiskUsage()
	ctx_error_handler(ctx, err)

	json_response, err := Json(result)
	ctx_error_handler(ctx, err)
	fmt.Fprintf(ctx, json_response)
}
