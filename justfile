default:
    @just --list --unsorted --justfile {{ justfile() }} | grep -v default

vet *args:
    go vet {{ args }} ./...

staticcheck *args:
    if [ ! -f $(go env GOPATH)/bin/staticcheck ]; then \
      go install honnef.co/go/tools/cmd/staticcheck@latest; \
    fi
    staticcheck {{ args }} ./...

gosec *args:
    if [ ! -f $(go env GOPATH)/bin/gosec ]; then \
      go install github.com/securego/gosec/v2/cmd/gosec@latest; \
    fi
    gosec -exclude-dir=testdata -terse {{ args }} ./...

lint: vet staticcheck gosec

fmt *args:
    go fmt {{ args }} ./...

build *args:
    go build {{ args }} .

install *args:
    go install {{ args }} .

test *args: install
    cd e2e && go clean -testcache && go test -v {{ args }} ./...

clean *args:
    go clean {{args}}

