package main

import "github.com/gofp/pkg/decorator"

func main() {

	g := decorator.GrpcHandler{Grpc_endpoint: "grpc://ep-1"}
	h := decorator.HttpHandler{Http_endpoint: "http://ep-1"}

	// Handler
	g.HandleMe("MyRequest")
	h.HandleMe("MyRequest")

	// Lets use decorated versions of these handlers
	// Tracer -> Handler
	tg := decorator.Tracer{Handler: &g, Trace_app_id: "TRACE-xx123"}
	tg.HandleMe("DecoratedMyRequest")

	// Lets use decorated versions of the "trace" decoratar
	// logger -> Tracer -> Handler
	ltg := decorator.Logger{Handler: &tg, Log_app_id: "LOG-xxx123"}
	ltg.HandleMe("FurtherDecoratedMyRequest")
}
