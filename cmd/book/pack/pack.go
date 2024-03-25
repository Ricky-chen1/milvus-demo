package pack

import (
	"errors"

	"github.com/Ricky-chen1/milvus-demo/kitex_gen/book"
	"github.com/Ricky-chen1/milvus-demo/pkg/errno"
)

func baseResp(err errno.ErrNo) *book.BaseResp {
	return &book.BaseResp{
		Code: err.ErrorCode,
		Msg:  &err.ErrorMsg,
	}
}

func BuildBaseResp(err error) *book.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceError.WithMessage(err.Error())
	return baseResp(s)
}
