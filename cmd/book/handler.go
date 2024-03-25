package main

import (
	"context"

	"github.com/Ricky-chen1/milvus-demo/cmd/book/pack"
	"github.com/Ricky-chen1/milvus-demo/cmd/book/service"
	book "github.com/Ricky-chen1/milvus-demo/kitex_gen/book"
	"github.com/Ricky-chen1/milvus-demo/pkg/errno"
)

// BookServiceImpl implements the last service interface defined in the IDL.
type BookServiceImpl struct{}

// Search implements the BookServiceImpl interface.
func (s *BookServiceImpl) Search(ctx context.Context, req *book.SearchReq) (resp *book.SearchResp, err error) {
	// TODO: Your code here...
	resp = new(book.SearchResp)

	idList, err := service.NewBookService(ctx).SearchWithVector(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.BookIdList = idList
	return resp, nil
}

// InsertData implements the BookServiceImpl interface.
func (s *BookServiceImpl) InsertData(ctx context.Context, req *book.InsertDataReq) (resp *book.InsertDataResp, err error) {
	// TODO: Your code here...
	resp = new(book.InsertDataResp)

	err = service.NewBookService(ctx).InsertData(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
