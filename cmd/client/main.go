package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os/signal"
	"syscall"

	"github.com/adrianoff/go-system-monitoring/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	conn, err := grpc.DialContext(ctx, "127.0.0.1:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewMonitoringServiceClient(conn)

	stream, err := client.StreamSnapshots(ctx, &pb.SnapshotRequest{
		WarmingUpTime:  uint32(15),
		SnapshotPeriod: uint32(5),
	})
	if err != nil {
		log.Fatal(err)
	}

	Label:
		for {
			select {
			case <-ctx.Done():
				break Label
			default:
				snapshot, err := stream.Recv()
				if errors.Is(err, io.EOF) {
					break Label
				}
				if err != nil {
					log.Fatal(err)
				}

				ProcessSnapshot(snapshot)
			}
		}
}

func ProcessSnapshot(snapshot *pb.Snapshot) {
	fmt.Printf("\nLoad average: %.02f %.02f %.02f", snapshot.LoadAverage.Min, snapshot.LoadAverage.Five, snapshot.LoadAverage.Fifteen)
}
