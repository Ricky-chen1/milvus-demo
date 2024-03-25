package milvus

import (
	"context"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
)

type Client struct {
	milvus client.Client
}

var MClient Client

func Init() {
	//...other snippet ...
	ctx := context.Background()
	client, err := client.NewClient(ctx, client.Config{
		Address: "0.0.0.0:19530",
	})
	if err != nil {
		panic(err)
		// handle error
	}
	m := Client{
		milvus: client,
	}

	MClient = m
}
