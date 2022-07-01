package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ellemouton/playgrounds/grpcerrors"
	"github.com/ellemouton/playgrounds/grpcerrors/pb"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "grpc-errors-cli"
	app.Flags = []cli.Flag{
		cli.Int64Flag{
			Name:  "serverport",
			Usage: "the port of the grpc server",
			Value: 8080,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "noerror",
			Action: noerror,
		},
		{
			Name:   "statusok",
			Action: statusok,
		},
		{
			Name:   "nonstatuserror",
			Action: nonstatuserror,
		},
		{
			Name:   "statuserror",
			Action: statuserror,
		},
		{
			Name:   "detailedstatuserror",
			Action: detailedstatuserror,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

type ErrCode uint8

const (
	ErrInternal     ErrCode = 0
	ErrRuleViolated ErrCode = 1
)

func detailedstatuserror(ctx *cli.Context) error {
	client, close, err := getClient(ctx)
	if err != nil {
		return err
	}
	defer close()

	_, err = client.StatusErrorWithDetails(
		context.Background(), &pb.Empty{},
	)
	fmt.Printf("error type: %T, error value: %v\n", err, err)

	s, ok := status.FromError(err)
	fmt.Printf("found status? %v, status values: %v\n", ok, s)

	for _, d := range s.Details() {
		switch err := d.(type) {
		case *pb.Error:
			switch e := err.Error.(type) {
			case *pb.Error_RuleViolationErr:
				fmt.Printf("%T: %v\n", e, e)
			case *pb.Error_InternalErr:
				fmt.Printf("%T: %v\n", e, e)
			}
		}
	}

	fmt.Printf("is ErrInternal? %v\n", errorIs(err, ErrInternal))
	fmt.Printf("is ErrRuleViolation? %v\n", errorIs(err, ErrRuleViolated))

	return nil
}

func errorIs(err error, code ErrCode) bool {
	s, ok := status.FromError(err)
	if !ok {
		return false
	}

	for _, d := range s.Details() {
		switch err := d.(type) {
		case *pb.Error:
			switch err.Error.(type) {
			case *pb.Error_RuleViolationErr:
				return code == ErrRuleViolated
			case *pb.Error_InternalErr:
				return code == ErrInternal
			}
		}
	}

	return false
}

func statuserror(ctx *cli.Context) error {
	client, close, err := getClient(ctx)
	if err != nil {
		return err
	}
	defer close()

	_, err = client.StatusError(context.Background(), &pb.Empty{})
	fmt.Printf("error type: %T, error value: %v\n", err, err)

	s, ok := status.FromError(err)
	fmt.Printf("found status? %v, status values: %v\n", ok, s)

	fmt.Printf("matches on custom error? %v\n",
		errors.Is(err, grpcerrors.ErrCustomError))
	return nil
}

func nonstatuserror(ctx *cli.Context) error {
	client, close, err := getClient(ctx)
	if err != nil {
		return err
	}
	defer close()

	_, err = client.NonStatusError(context.Background(), &pb.Empty{})
	fmt.Printf("error type: %T, error value: %v\n", err, err)

	s, ok := status.FromError(err)
	fmt.Printf("found status? %v, status values: %v\n", ok, s)

	fmt.Printf("matches on custom error? %v\n",
		errors.Is(err, grpcerrors.ErrCustomError))
	return nil
}

func noerror(ctx *cli.Context) error {
	client, close, err := getClient(ctx)
	if err != nil {
		return err
	}
	defer close()

	_, err = client.NoError(context.Background(), &pb.Empty{})
	fmt.Printf("error type: %T, error value: %v\n", err, err)

	s, ok := status.FromError(err)
	fmt.Printf("found status? %v, status values: %v\n", ok, s)

	return nil
}

func statusok(ctx *cli.Context) error {
	client, close, err := getClient(ctx)
	if err != nil {
		return err
	}
	defer close()

	_, err = client.StatusOk(context.Background(), &pb.Empty{})
	fmt.Printf("error type: %T, error value: %v\n", err, err)

	s, ok := status.FromError(err)
	fmt.Printf("found status? %v, status values: %v\n", ok, s)

	return nil
}

func getClient(ctx *cli.Context) (pb.ErrorsClient, func(), error) {
	addr := fmt.Sprintf("localhost:%d", ctx.GlobalInt64("serverport"))
	fmt.Printf("dialing %s\n", addr)

	grpcConn, err := grpc.DialContext(
		context.Background(), addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, err
	}

	return pb.NewErrorsClient(grpcConn), func() {
		if err := grpcConn.Close(); err != nil {
			log.Fatalln(err)
		}
	}, nil
}
