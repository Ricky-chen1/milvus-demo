## milvus-demo

### 运行

在**根目录**下,构建相关环境
```bash
    make env-up
```

在**cmd/book目录**下，编译生成二进制文件并执行
```bash
    make server
```

最后在**cmd/client**目录下,启动客户端并发送rpc请求进行milvus搜索
```bash
    go run main.go
```

### make指令列表

```bash
    # cmd/book下
    make gen        # 更新kitex生成代码
    make server     # 启动kitex服务端

    # 根目录下
    make env-up     # 构建相关环境
    make env-down   # 结束环境
    make clean      # 环境结束后删除数据卷
```