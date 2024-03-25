package service

import (
	"github.com/Ricky-chen1/milvus-demo/cmd/book/dal/milvus"
	"github.com/Ricky-chen1/milvus-demo/kitex_gen/book"
)

func (bs *BookService) SearchWithVector(req *book.SearchReq) ([]int64, error) {
	idList, err := milvus.Search(bs.ctx, req)
	if err != nil {
		return nil, err
	}

	return idList, nil
}
