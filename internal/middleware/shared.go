package middleware

import "google.golang.org/grpc"

var UnaryServerMiddleware []grpc.UnaryServerInterceptor = []grpc.UnaryServerInterceptor{
	mimeTypeCheckInterceptor,
	mimeMismatchInterceptor,
}
