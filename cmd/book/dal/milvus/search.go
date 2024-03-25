package milvus

import (
	"context"

	"github.com/Ricky-chen1/milvus-demo/kitex_gen/book"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

func Search(ctx context.Context, req *book.SearchReq) ([]int64, error) {
	m := MClient.milvus

	idx, err := entity.NewIndexIvfFlat( // NewIndex func
		entity.L2, // metricType
		1024,      // ConstructParams
	)
	if err != nil {
		return nil, err
	}

	if err = m.CreateIndex(
		ctx,                // ctx
		req.CollectionName, // CollectionName
		"book_intro",       // fieldName
		idx,                // entity.Index
		false,              // async
	); err != nil {
		return nil, err
	}

	err = m.LoadCollection(ctx, req.CollectionName, false)
	if err != nil {
		return nil, err
	}

	sp, _ := entity.NewIndexIvfFlatSearchParam( // NewIndex*SearchParam func
		10, // searchParam
	)

	opt := client.SearchQueryOptionFunc(func(option *client.SearchQueryOption) {
		option.Limit = 3
		option.Offset = 0
		option.ConsistencyLevel = entity.ClStrong //
		option.IgnoreGrowing = false
	})

	floatv := make([]float32, 0, req.Dim)
	for _, f64 := range req.Vector {
		floatv = append(floatv, float32(f64))
	}

	searchResult, err := m.Search(
		ctx,                 // ctx
		req.CollectionName,  // CollectionName
		[]string{},          // partitionNames
		"",                  // expr
		[]string{"book_id"}, // outputFields
		[]entity.Vector{entity.FloatVector(floatv)}, // vectors
		"book_intro", // vectorField
		entity.L2,    // metricType
		10,           // topK
		sp,           // searchParams
		opt,
	)
	if err != nil {
		return nil, err
	}

	// TODO: return result
	bookIDList := make([]int64, 0, req.ResultCount)

	for _, sr := range searchResult {
		for num := 0; num < sr.ResultCount; num++ {
			id, err := sr.IDs.GetAsInt64(num)
			if err != nil {
				continue
			}
			bookIDList = append(bookIDList, id)
		}
	}

	err = m.ReleaseCollection(ctx, req.CollectionName)
	if err != nil {
		return nil, err
	}

	return bookIDList, nil
}
