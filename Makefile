build:
	go build -o ./bin/welp ./cmd/welp
test: 
	go test ./cmd/welp
run:
	air --build.cmd "go build -o bin/welp cmd/welp/main.go" --build.bin "./bin/welp" --build.exclude_dir "test,tmp,bin"
