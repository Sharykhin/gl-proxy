serve:
	HTTP_ADDRESS=127.0.0.1:8888 CORS_ORIGIN=* go run main.go

docker-serve:
	go run main.go

lint:
	gometalinter ./...

test:
	go test -v ./...