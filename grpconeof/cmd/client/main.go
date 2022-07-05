package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/ellemouton/playgrounds/grpconeof/pb_old"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	app := cli.NewApp()
	app.Name = "grpc-oneof-cli"
	app.Flags = []cli.Flag{
		cli.Int64Flag{
			Name:  "serverport",
			Usage: "the port of the grpc server",
			Value: 8080,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "old",
			Action: newOrOld(false),
		},
		{
			Name:   "new",
			Action: newOrOld(true),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

func newOrOld(new bool) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		client, close, err := getClient(ctx)
		if err != nil {
			return err
		}
		defer close()

		resp, err := client.ListItems(
			context.Background(), &pb.ListItemsRequest{
				New: new,
			},
		)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("no error!")
		fmt.Printf("%+v\n", resp)

		switch resp.Messages.(type) {
		case *pb.ListItemsResponse_OldMsg:
			fmt.Println("type is `OldMsg`")
		default:
			fmt.Println("unknown type!")
		}
		return nil
	}
}

func getClient(ctx *cli.Context) (pb.CalendarClient, func(), error) {
	addr := fmt.Sprintf("localhost:%d", ctx.GlobalInt64("serverport"))
	fmt.Printf("dialing %s\n\n", addr)

	grpcConn, err := grpc.DialContext(
		context.Background(), addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, err
	}

	return pb.NewCalendarClient(grpcConn), func() {
		if err := grpcConn.Close(); err != nil {
			log.Fatalln(err)
		}
	}, nil
}
