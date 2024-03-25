###client

```
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=client --go-grpc_out=client proto/blog.proto
```

###server

```export PATH="$PATH:$(go env GOPATH)/bin"
   protoc --go_out=. --go-grpc_out=. proto/blog.proto
```

###Reference

```https://sahansera.dev/building-grpc-server-go/```
