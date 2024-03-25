package main

import (
	"log"

	"github.com/Ricky-chen1/milvus-demo/cmd/book/dal"
	book "github.com/Ricky-chen1/milvus-demo/kitex_gen/book/bookservice"
	"github.com/Ricky-chen1/milvus-demo/pkg/constants"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)

func main() {
	dal.Init()

	svr := book.NewServer(
		new(BookServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.BookServiceName,
		}))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
