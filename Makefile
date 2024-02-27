compile_deprecated:
	protoc -I=. --go_out=./proto ./studentpb/*.proto

compile:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_out=. ./studentpb/*.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_out=. ./testpb/*.proto
