# traileau
A clone of Trello build in Go 


generate the proto file
```shell
# at the root of the microservice
protoc --go_out=. --go-grpc_out=. proto/auth/auth.proto
```
