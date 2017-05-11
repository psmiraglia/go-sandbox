BUILD=`date +%Y/%m/%dT%H:%M:%S`
COMMIT=`git rev-parse HEAD | cut -c1-8`
LDFLAGS=-ldflags "-X github.com/psmiraglia/go-sandbox/sandbox/common.Build=${BUILD} -X github.com/psmiraglia/go-sandbox/sandbox/common.Commit=${COMMIT}"

run:
	go run ${LDFLAGS} main.go

build:
	go build ${LDFLAGS}
