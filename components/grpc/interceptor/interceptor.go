package interceptor

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

type ContextKey struct {
	name string
}

func (k *ContextKey) String() string {
	return "grpc context value " + k.name
}

func WithUnaryClientChain(interceptors ...grpc.UnaryClientInterceptor) grpc.DialOption {
	return grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(interceptors...))
}

func WithStreamClientChain(interceptors ...grpc.StreamClientInterceptor) grpc.DialOption {
	return grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(interceptors...))
}
