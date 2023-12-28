package decorator

import "fmt"

// Want to augment an object with additional functionality
// Do not want to rewrite or alter existing code (OCP principle)
// want to keep new features separate (SRP principle)
// Cannot break existing code

// Goal: Embed the *new* object in the older object and extend the functionality without breaking existing code.

// Design a logger that is traced for communication handlers.

type CommsHandler interface {
	HandleMe(request string) string
}

type GrpcHandler struct {
	Grpc_endpoint string
}

func (g *GrpcHandler) HandleMe(req string) string {
	fmt.Printf("   %s request received - %s\n", "Grpc", req)
	return req
}

type HttpHandler struct {
	Http_endpoint string
}

func (h *HttpHandler) HandleMe(req string) string {
	fmt.Printf("   %s request received - %s\n", "Http", req)
	return req
}

type Tracer struct {
	Handler      CommsHandler
	Trace_app_id string
}

func (t *Tracer) HandleMe(req string) string {
	fmt.Printf("  %s-%s request received - %s\n", "Trace", t.Trace_app_id, t.Handler.HandleMe(req))
	return req
}

type Logger struct {
	Handler    CommsHandler
	Log_app_id string
}

func (l *Logger) HandleMe(req string) string {
	fmt.Printf("%s-%s request received - %s\n", "Log", l.Log_app_id, l.Handler.HandleMe(req))
	return req
}
