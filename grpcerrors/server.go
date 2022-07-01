package grpcerrors

import (
	"context"
	"fmt"

	"github.com/ellemouton/playgrounds/grpcerrors/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrCustomError = fmt.Errorf("custom error to match on")
)

type Server struct {
	pb.UnimplementedErrorsServer
}

func (*Server) NoError(_ context.Context, _ *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (*Server) StatusOk(_ context.Context, _ *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, status.New(codes.OK, "all groovy").Err()
}

func (*Server) NonStatusError(_ context.Context, _ *pb.Empty) (*pb.Empty,
	error) {

	return nil, ErrCustomError
}

func (*Server) StatusError(_ context.Context, _ *pb.Empty) (*pb.Empty, error) {
	return nil, status.New(
		codes.ResourceExhausted, ErrCustomError.Error(),
	).Err()
}

func (*Server) StatusErrorWithDetails(_ context.Context,
	_ *pb.Empty) (*pb.Empty, error) {

	st := status.New(codes.ResourceExhausted, "something something")
	st, err := st.WithDetails(&pb.Error{
		Error: &pb.Error_RuleViolationErr{
			RuleViolationErr: &pb.ErrRuleViolation{
				RuleName: "rate-limit",
				Err:      "read rate-limit violated",
			},
		},
	})
	if err != nil {
		panic(err)
	}

	return nil, st.Err()
}
