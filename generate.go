package forge

// format .proto before generating from it
//go:generate buf format -w

// generate .go from .proto
//go:generate buf generate .

// format .go, including .go generated by .proto
// were it to not meet gofmt's standards
//go:generate go fmt ./...

//go:generate golangci-lint run --fix

// build the shim binary which is often copied into containers
// and used as their entrypoint
//go:generate env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./internal/bin/shim ./internal/cmd/shim

// package up the shim binary so that copy times are as fast as possible
//go:generate upx --ultra-brute ./internal/bin/shim

// tidy up
//go:generate go mod tidy
