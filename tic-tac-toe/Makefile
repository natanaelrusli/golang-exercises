run:
	go run ./cmd/tictactoe

test:
	go test ./... -v

coverage:
	go test ./... -v -coverprofile=cover.out
	go tool cover -html=cover.out