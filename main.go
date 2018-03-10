package main

import (
	"github.com/goinggo/tracelog"
	"net"
	"os"
	"google.golang.org/grpc"
	pb_email "github.com/onezerobinary/email-box/proto"
	"github.com/onezerobinary/email-box/email"
	"github.com/spf13/viper"
)

const (
	GRPC_PORT = ":1976"
)

func main(){

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		tracelog.Errorf(err, "main", "main", "Error reading config file")
	}

	listen, err := net.Listen("tcp", GRPC_PORT)
	if err != nil {
		tracelog.Errorf(err, "app", "main", "Failed to start the service")
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	// Add to the grpcServer the Service
	pb_email.RegisterEmailServiceServer(grpcServer, &email.EmailServiceServer{})

	tracelog.Trace("app", "main", "Grpc Server Listening on port 1976")

	grpcServer.Serve(listen)

}
