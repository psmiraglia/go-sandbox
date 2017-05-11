# My Go Sandbox

It's time to add Go language to my skillset...

## Play with my sandbox

    $ go get github.com/psmiraglia/go-sandbox/sandbox
    $ cd $GOPATH/src/github.com/psmiraglia/go-sandbox/sandbox
    $ make build
    $ ./sanbox doit
    {"@build":"2017/05/11T18:26:30","@commit":"5b7488ce","@timestamp":"2017-05-11T18:37:58.147332764+02:00","@version":1,"level":"debug","msg":"Hello! This is Debug..."}
    {"@build":"2017/05/11T18:26:30","@commit":"5b7488ce","@timestamp":"2017-05-11T18:37:58.147543661+02:00","@version":1,"level":"info","msg":"Hello! This is Info..."}
    {"@build":"2017/05/11T18:26:30","@commit":"5b7488ce","@timestamp":"2017-05-11T18:37:58.147574451+02:00","@version":1,"level":"warning","msg":"Hello! This is Warning..."}
    {"@build":"2017/05/11T18:26:30","@commit":"5b7488ce","@timestamp":"2017-05-11T18:37:58.147603414+02:00","@version":1,"level":"error","msg":"Hello! This is Error..."}
    {"@build":"2017/05/11T18:26:30","@commit":"5b7488ce","@timestamp":"2017-05-11T18:37:58.147625635+02:00","@version":1,"level":"fatal","msg":"Hello! This is Fatal..."}

## References

*   [Go Language](https://golang.org/)
*   [sirupsen/logrus](https://github.com/sirupsen/logrus)
*   [urfave/cli](https://github.com/urfave/cli)
