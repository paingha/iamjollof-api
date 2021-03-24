package dialer

import (
	"bitbucket.com/iamjollof/server/plugins"
	"google.golang.org/grpc"
)

//Dial - connects to gRPC serverand returns the connection
func Dial(port string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		plugins.LogFatal("gRPC Server internal Client", "did not connect", err)
		return nil, err
	}
	defer conn.Close()
	return conn, nil
}
