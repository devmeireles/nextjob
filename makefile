test-verbose:
	go test ./tests -v

test-controllers:
	go test -v ./controller -coverprofile=coverage.out

see-converage:
	go tool cover -html=coverage.out

build-swagger:
	swag init