GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

dev:
	go run main.go

test:
	go test -cover ./...

build:
	go build main.go

start-service:
	docker rm keyforge-name-of-the-day
	docker build . -t keyforge-name-of-the-day
	docker run -d --name keyforge-name-of-the-day \
		-p 5001:5001 \
		--restart=always \
		keyforge-name-of-the-day

stop-service:
	docker stop keyforge-name-of-the-day
	docker rm keyforge-name-of-the-day
