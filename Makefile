fmt:
	go fmt ./...

test:
	go test -v -cover ./...

test_cov:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

test_cov_ci:
	go test ./... -coverprofile=coverage.out

build:
	go build