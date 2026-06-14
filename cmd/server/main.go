package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/khazeez/user-service/proto"

	grpcHandler "github.com/khazeez/user-service/internal/transport/grpc"

	"github.com/khazeez/user-service/internal/repository"
	"github.com/khazeez/user-service/internal/service"
)

func main() {

    repo := repository.NewMemoryRepo()

    userService := service.NewUserService(
        repo,
    )

    handler := grpcHandler.NewUserHandler(
        userService,
    )

    lis, err := net.Listen(
        "tcp",
        ":50051",
    )

    if err != nil {
        log.Fatal(err)
    }

    server := grpc.NewServer()

    pb.RegisterUserServiceServer(
        server,
        handler,
    )

    log.Println(
        "grpc server running :50051",
    )

    server.Serve(lis)
}