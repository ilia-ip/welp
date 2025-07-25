build: templ
	go build -o ./bin/welp ./cmd/welp
test: 
	go test ./cmd/welp
run: templ
	go run ./cmd/welp
templ:
	go tool templ generate ./internal/frontend/pages/
