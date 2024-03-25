package milvus

import (
	"context"
	"math/rand"

	"github.com/Ricky-chen1/milvus-demo/kitex_gen/book"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

func CreateCollection(ctx context.Context, req *book.InsertDataReq) error {
	m := MClient.milvus

	exist, err := m.HasCollection(ctx, req.CollectionName)
	if err != nil {
		return err
	}

	if exist {
		// colletion exist,不作为错误
		return nil
	}

	schema := &entity.Schema{
		CollectionName: req.CollectionName,
		Description:    "search for book",
		Fields: []*entity.Field{
			{
				Name:       "book_id",
				DataType:   entity.FieldTypeInt64,
				PrimaryKey: true,
				AutoID:     false,
			},
			{
				Name:       "word_count",
				DataType:   entity.FieldTypeInt64,
				PrimaryKey: false,
				AutoID:     false,
			},
			{
				Name:     "book_intro",
				DataType: entity.FieldTypeFloatVector,
				TypeParams: map[string]string{
					"dim": "2",
				},
			},
		},
		EnableDynamicField: true,
	}

	if err := m.CreateCollection(ctx, schema, 2); err != nil {
		return err
	}

	return nil
}

func InsertData(ctx context.Context, req *book.InsertDataReq) error {
	m := MClient.milvus

	bookIDs := make([]int64, 0, req.DataCount)
	wordCounts := make([]int64, 0, req.DataCount)
	bookIntros := make([][]float32, 0, req.DataCount)
	for i := 0; i < int(req.DataCount); i++ {
		bookIDs = append(bookIDs, int64(i))
		wordCounts = append(wordCounts, int64(i+10000))
		v := make([]float32, 0, 2)
		for j := 0; j < 2; j++ {
			v = append(v, rand.Float32())
		}
		bookIntros = append(bookIntros, v)
	}
	idColumn := entity.NewColumnInt64("book_id", bookIDs)
	wordColumn := entity.NewColumnInt64("word_count", wordCounts)
	introColumn := entity.NewColumnFloatVector("book_intro", 2, bookIntros)

	if _, err := m.Insert(
		ctx,
		req.CollectionName,
		"",
		idColumn,
		wordColumn,
		introColumn,
	); err != nil {
		return err
	}

	return nil
}
