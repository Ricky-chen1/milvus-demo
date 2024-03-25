namespace go book

struct BaseResp {
    1: required i64 code,
    2: optional string msg,
}

struct Book{
    1: required i64 id
    2: required i64 word_count
}

struct InsertDataReq{
    1: required string collection_name
    2: required i64 data_count
}

struct InsertDataResp{
    1: required BaseResp base
}

struct SearchReq{
    1: required string collection_name
    2: required i64 dim // 向量维数
    3: required list<double> vector
    4: required i64 result_count 
}

struct SearchResp{
    1: required BaseResp base
    2: optional list<i64> book_id_list
}

service BookService{
    SearchResp Search(1: SearchReq req)
    InsertDataResp InsertData(1: InsertDataReq req)
}