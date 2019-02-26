default:
	GOOS=linux GOARCH=amd64 go build -o gravity-grpc-sidecar-filter-example.linux
	GOOS=darwin GOARCH=amd64 go build -o gravity-grpc-sidecar-filter-example.darwin