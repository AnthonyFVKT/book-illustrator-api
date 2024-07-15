package service

import (
	"context"
	"github.com/AnthonyFVKT/book-illustrator-api/internal/rpc"
	pb "github.com/AnthonyFVKT/book-illustrator-srv/proto/illustrator"
)

type Illustrator struct {
	rep *rpc.Illustrator
}

func NewIllustrator(rep *rpc.Illustrator) *Illustrator {
	return &Illustrator{
		rep: rep,
	}
}

func (s *Illustrator) Create(ctx context.Context, text string) (*pb.CreateResponse, error) {
	return s.rep.Create(ctx, text)
}
