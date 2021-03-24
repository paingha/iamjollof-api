package interceptors

import (
	"context"
	"strings"

	"bitbucket.com/iamjollof/server/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

//Pass in accessToken and then if the route is not public verify the access code
//and then verify the access code authorization.
func verifyRouteAuth(route string, routes map[string]string, accessToken string) bool {
	for key, element := range routes {
		if route == key && element == "public" {
			return true
		}
		if route == key && element == "admin" {
			claims, correct := security.VerifyJWT(accessToken)
			if !correct || !claims.IsAdmin {
				return false
			}
			return true
		}
	}
	return false
}

//AuthUnary - Unary Authentication interceptor
//where headers is md; which is a map[string][]string
func AuthUnary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		headers, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
		}
		values := headers["authorization"]
		if len(values) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "authentication token is not provided")
		}
		accessToken := strings.Join(values, "")
		if authorized := verifyRouteAuth(info.FullMethod, Routes, accessToken); authorized {
			return handler(ctx, req)
		}
		return nil, status.Errorf(codes.PermissionDenied, "Not Authorized to access resource")
	}
}

//AuthStream - Stream Authentication interceptor
//where headers is md; which is a map[string][]string
func AuthStream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		headers, ok := metadata.FromIncomingContext(stream.Context())
		if !ok {
			return status.Errorf(codes.Unauthenticated, "metadata is not provided")
		}

		values := headers["authorization"]
		if len(values) == 0 {
			return status.Errorf(codes.Unauthenticated, "authentication token is not provided")
		}
		accessToken := strings.Join(values, "")
		if authorized := verifyRouteAuth(info.FullMethod, Routes, accessToken); authorized {
			return handler(srv, stream)
		}
		return status.Errorf(codes.PermissionDenied, "Not Authorized to access resource")
	}
}
