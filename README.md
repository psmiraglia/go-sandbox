# My Go Sandbox

It's time to add Go language to my skillset...

## Play with my sandbox

    $ go get github.com/psmiraglia/go-sandbox/sandbox
    $ cd $GOPATH/src/github.com/psmiraglia/go-sandbox/sandbox
    $ make build
    $ ./sanbox doit | jq .
    {
      "@build": "2017/05/11T19:28:54",
      "@commit": "b47da0d8",
      "@timestamp": "2017-05-11T19:32:00.144495616+02:00",
      "@version": 1,
      "hash": "ZxCW7M0f+MNGKiXmfSp6lSv3EZML+svZYdCxeHpV2Bg=",
      "level": "debug",
      "msg": "Hello! This is Debug..."
    }
    {
      "@build": "2017/05/11T19:28:54",
      "@commit": "b47da0d8",
      "@timestamp": "2017-05-11T19:32:00.144738068+02:00",
      "@version": 1,
      "hash": "CkufJm5GQ2vMCMbLDx+8Wg5kzF7z3j4fLSbNV+JqoCs=",
      "level": "info",
      "msg": "Hello! This is Info..."
    }
    {
      "@build": "2017/05/11T19:28:54",
      "@commit": "b47da0d8",
      "@timestamp": "2017-05-11T19:32:00.144781964+02:00",
      "@version": 1,
      "hash": "OQdtGVnRUxYpmLSN8nBCYCjgmtR/rIVH2GYtSYEA6m4=",
      "level": "warning",
      "msg": "Hello! This is Warning..."
    }
    {
      "@build": "2017/05/11T19:28:54",
      "@commit": "b47da0d8",
      "@timestamp": "2017-05-11T19:32:00.144823008+02:00",
      "@version": 1,
      "hash": "5x7m2yWdYX3+UaqlPbM9JM8m0jGKuoBXl9PIGzvkaRI=",
      "level": "error",
      "msg": "Hello! This is Error..."
    }
    {
      "@build": "2017/05/11T19:28:54",
      "@commit": "b47da0d8",
      "@timestamp": "2017-05-11T19:32:00.144859314+02:00",
      "@version": 1,
      "hash": "fNqYxPUfbYSvtps40jJAntroCTHgRhvHe+GtKF4njdo=",
      "level": "fatal",
      "msg": "Hello! This is Fatal..."
    }

## References

*   [Go Language](https://golang.org/)
*   [sirupsen/logrus](https://github.com/sirupsen/logrus)
*   [spf13/viper](https://github.com/spf13/viper)
*   [urfave/cli](https://github.com/urfave/cli)
