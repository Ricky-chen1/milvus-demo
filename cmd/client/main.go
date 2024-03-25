package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Ricky-chen1/milvus-demo/kitex_gen/book"
	"github.com/Ricky-chen1/milvus-demo/kitex_gen/book/bookservice"
	"github.com/Ricky-chen1/milvus-demo/pkg/constants"
	"github.com/cloudwego/kitex/client"
)

func main() {

	client, err := bookservice.NewClient(
		constants.BookServiceName,
		client.WithHostPorts("0.0.0.0:8888"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ireq := &book.InsertDataReq{
		CollectionName: "book",
		DataCount:      1000,
	}

	_, err = client.InsertData(ctx, ireq)
	if err != nil {
		log.Fatal(err)
	}

	vector := []float64{0.1, 0.2}

	sreq := &book.SearchReq{
		CollectionName: "book",
		ResultCount:    1000,
		Vector:         vector,
		Dim:            2,
	}

	sresp, err := client.Search(ctx, sreq)
	if err != nil {
		log.Fatal(err)
	}

	// 与输入向量最相似的bookIDList
	for _, bookID := range sresp.BookIdList {
		fmt.Printf("id: %v ", bookID)
	}
}
