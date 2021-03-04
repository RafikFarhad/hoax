# Hoax

A basic boilerplate for developing Rest API using GoLang and Fiber

### Ingredients
- [x] [fiber](https://github.com/gofiber/fiber)
- [x] mysql-driver
- [x] [ORM](https://github.com/go-gorm/gorm)
- [x] [INI config parser](https://gopkg.in/ini.v1)
- [ ] authentication
- [ ] template engine (!)
- [ ] swagger-ui
- [ ] web-socket support
- [ ] docker image

### Usage
- Clone the repo
- `go mod download`
- `go run main.go`

There is a `config.sample.ini` file which will be parsed by default. Custom `config.ini` file can be provided by
```shell
go run main.go -c config.ini
```

#### Other Args
```shell
$ go build ./main.go
$ ./main--help
  -c string
        Config File (default "config.sample.ini")
  -h string
        Host (default "127.0.0.1")
  -p int
        Port (default 3000)

```

### License
Not decided yet