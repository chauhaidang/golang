package main

import (
	"context"
	"testing"

	"dang.com/stub/service"
	"go.nhat.io/grpcmock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCDownStreamCall(t *testing.T) {
	s := grpcmock.NewServer(
		grpcmock.WithPort(4000),
		grpcmock.RegisterService(service.RegisterMyServiceServer),
		func(s *grpcmock.Server) {
			s.
				ExpectUnary("service.MyService/Get").
				WithPayload(&service.Request{Name: "ABCD"}).
				ReturnJSON(map[string]string{"foo": "bar"})
		},
	)
	defer s.Close()

	conn, err := grpc.Dial(s.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		t.Fatal(err)
	}
	client := service.NewMyServiceClient(conn)

	client.Get(context.Background(), &service.Request{Name: "ABC"})
	err = s.ExpectationsWereMet()
	if err != nil {
		t.Error("DOWNSTREAM IS NOT CALLED!")
	}
}
