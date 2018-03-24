serve:
	HTTP_ADDRESS=127.0.0.1:8888 go run main.go

lint:
	gometalinter ./...

test:
	go test -v ./...