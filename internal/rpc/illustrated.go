package rpc

import (
	"context"
	pb "github.com/AnthonyFVKT/book-illustrator-srv/proto/illustrator"
)

type Illustrator struct {
	illustratorClient pb.IllustratorServiceClient
}

func NewIllustrator(illustratorClient pb.IllustratorServiceClient) *Illustrator {
	return &Illustrator{illustratorClient: illustratorClient}
}

func (r *Illustrator) Create(ctx context.Context, text string) (*pb.CreateResponse, error) {
	return r.illustratorClient.Create(ctx, &pb.CreateRequest{Text: text})
}
