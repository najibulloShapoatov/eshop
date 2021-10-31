build:
	docker-compose build app

run:
	docker-compose up app

test:
	go test -v ./...

swag:
	swag init -g cmd/main.go
#$GOPATH/bin/swag init -g cmd/main.go
	
run-go:
	go run cmd/main.go
