BUILD=`date +%Y/%m/%dT%H:%M:%S`
COMMIT=`git rev-parse HEAD`
LDFLAGS=-ldflags "-X github.com/psmiraglia/go-sandbox/sandbox/common.build=${BUILD} -X github.com/psmiraglia/go-sandbox/sandbox/common.commit=${COMMIT}"

run:
	go run ${LDFLAGS} main.go

build:
	go build ${LDFLAGS}
