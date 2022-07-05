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

func (*Server) TestMap(_ context.Context, req *pb.TestMapRequest) (
	*pb.TestMapResponse, error) {

	if req.New {
		return &pb.TestMapResponse{
			Things: map[string]*pb.Thing{
				"new thing!": {
					Messages: &pb.Thing_NewMsg{
						NewMsg: &pb.NewMessage{
							Msg:   "new message!",
							Count: 5,
						},
					},
				},
				"old thing": {
					Messages: &pb.Thing_OldMsg{
						OldMsg: &pb.OldMessage{
							Msg: "old message",
						},
					},
				},
			},
		}, nil
	}

	return &pb.TestMapResponse{
		Things: map[string]*pb.Thing{
			"old thing": {
				Messages: &pb.Thing_OldMsg{
					OldMsg: &pb.OldMessage{
						Msg: "old message",
					},
				},
			},
		},
	}, nil
}
