BINARY=sandbox
BUILD=`date +%Y-%m-%dT%H:%M:%S`
COMMIT=`git rev-parse --short HEAD`
LDFLAGS=-ldflags "                                                           \
	-X github.com/psmiraglia/go-sandbox/sandbox.build=${BUILD}       \
	-X github.com/psmiraglia/go-sandbox/sandbox.commit=${COMMIT}     \
"

run:
	go run ${LDFLAGS} cmd/sandbox/main.go doit

build:
	go build ${LDFLAGS} -o ${BINARY} cmd/sandbox/main.go

install:
	go install ${LDFLAGS}
