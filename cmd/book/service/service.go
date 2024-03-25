package service

import (
	"context"

	"github.com/Ricky-chen1/milvus-demo/cmd/book/dal/milvus"
	"github.com/Ricky-chen1/milvus-demo/kitex_gen/book"
)

type BookService struct {
	ctx context.Context
}

func NewBookService(ctx context.Context) *BookService {
	return &BookService{ctx: ctx}
}

func (bs *BookService) InsertData(req *book.InsertDataReq) error {
	if err := milvus.CreateCollection(bs.ctx, req); err != nil {
		return err
	}

	if err := milvus.InsertData(bs.ctx, req); err != nil {
		return err
	}

	return nil
}
