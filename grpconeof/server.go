package grpconeof

import (
	"context"
	"fmt"

	pb "github.com/ellemouton/playgrounds/grpconeof/pb_new"
)

var (
	ErrCustomError = fmt.Errorf("custom error to match on")
)

type Server struct {
	pb.UnimplementedCalendarServer
}

func (*Server) ListItems(_ context.Context, req *pb.ListItemsRequest) (
	*pb.ListItemsResponse, error) {

	if req.New {
		return &pb.ListItemsResponse{
			Messages: &pb.ListItemsResponse_NewMsg{
				NewMsg: &pb.NewMessage{
					Msg:   "new message!",
					Count: 50,
				},
			},
		}, nil
	}

	return &pb.ListItemsResponse{
		Messages: &pb.ListItemsResponse_OldMsg{
			OldMsg: &pb.OldMessage{
				Msg: "old message",
			},
		},
	}, nil
}
