setup:
	cp internal/app/config.example.yaml internal/app/config.yaml && \
	go get ./... && \
	go install github.com/swaggo/swag/cmd/swag@latest

run:
	swag init && go run main.go