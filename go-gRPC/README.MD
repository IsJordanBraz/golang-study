protoc -I greet/proto --go_out=. --go-grpc_out=. greet/proto/dummy.proto
protoc -I greet/proto --go_out=. --go_opt=module=go-grpc --go-grpc_out=. --go-grpc_opt=module=go-grpc greet/proto/dummy.proto