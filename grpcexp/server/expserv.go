package server

import (
	"context"
	"errors"
	"fmt"
	"grpcexp/src/exp/proto/exp"
)

type ExpServ struct {
}

func (s *ExpServ) GetExp(ctx context.Context, in *exp.GetExpReq) (*exp.GetExpRes, error) {
	if in.Count <= 0 {
		return nil, errors.New("Out of range")
	}

	res := &exp.GetExpRes{
		Exps: make([]*exp.Exp, in.Count),
	}
	var i int64
	for i = 0; i < in.Count; i++ {
		res.Exps[i] = &exp.Exp{
			Id:   i,
			Name: fmt.Sprintf("%s, %d", "Guest", i),
		}
	}
	return res, nil
}

func (s *ExpServ) mustEmbedUnimplementedExpServiceServer() {}
