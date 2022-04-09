mock.gen:
	go generate -v ./...

test:
	go test -v -race ./...

build.server:
	GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o tts-server ./cmd/server/main.go

build.client:
	GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o say ./cmd/client/main.go

proto.gen:
	protoc --proto_path=infra/grpc infra/grpc/protofiles/*.proto --go_out=. --go-grpc_out=.
