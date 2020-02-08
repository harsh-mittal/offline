package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func rootHandler() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		log.Printf("Offline Request: %+v", ctx)

		log.Printf("Offline Request ID: %v\n", ctx.ID())
		log.Printf("Offline Request start time: %v\n", time.Now().Unix())
		time.Sleep(2*time.Minute + 10*time.Millisecond)
		log.Printf("Offline Request end ime: %v\n", time.Now().Unix())

		log.Printf("Offline Response: %+v", ctx)
	}
}

func main() {
	system := make(chan os.Signal, 2)

	address := fmt.Sprintf(":%d", 8010)
	handler := rootHandler()

	handlers := map[string]fasthttp.RequestHandler{
		"/": handler,
	}

	fastHandler := func(ctx *fasthttp.RequestCtx) {
		if h, ok := handlers[string(ctx.Path())]; ok {
			h(ctx)
		} else {
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	go func() {
		httpServer := &fasthttp.Server{
			Concurrency:  10240,
			Handler:      fastHandler,
			ReadTimeout:  500000000,
			WriteTimeout: 500000000,
		}

		if err := httpServer.ListenAndServe(address); err != nil {
			log.Printf("http server err=[%v]\n", err)
		}
		system <- syscall.SIGKILL
	}()
	log.Printf("QCMS server starts successfully, address: %s", address)

	signal.Notify(system, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)
	<-system

	log.Printf("QCMS server is shutting down, address: %s", address)

	close(system)

	return
}
