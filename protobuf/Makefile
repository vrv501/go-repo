export GO_MODULE = github.com/vrv501/go-repo/example-grpc
protogen:
	@protoc --proto_path=./proto --go_out=. --go-grpc_out=. --go_opt=module=${GO_MODULE} --go-grpc_opt=module=${GO_MODULE} proto/*.proto
.PHONY: protogen