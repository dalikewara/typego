# typego

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/dalikewara/typego)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dalikewara/typego)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/dalikewara/typego)
![GitHub license](https://img.shields.io/github/license/dalikewara/typego)

**typego** provides custom type that can be used to construct information (such as success data, error data, etc).

## Getting started

### Installation

You can use the `go get` method:

```bash
go get github.com/dalikewara/typego
```

### Usage

#### Error

`typego.Error` compatible with `error` interface, so you can use it as an `error` handler. `typego.Error`
has several methods that can be used to construct error information:

```go
type Error interface {
	ChangeCode(code string) Error
	ChangeMessage(message string) Error
	AddInfo(info ...interface{}) Error
	SetHttpStatus(httpStatus int) Error
	SetRPCStatus(rpcStatus int) Error
	GetCode() string
	GetMessage() string
	GetInfo() []string
	GetHttpStatus() int
	GetRPCStatus() int
	Error() string
}
```

and it will generate the error information based on this structure:

```go
type errorModel struct {
	Code       string   `json:"code"`
	Message    string   `json:"message"`
	Info       []string `json:"info"`
	HttpStatus int      `json:"http_status,omitempty"`
	RPCStatus  int      `json:"rpc_status,omitempty"`
}
```

For example:

```go
func main() {
    if err := myFunc(); err != nil {
        fmt.Println(err)
        // output
        // error: {"code":"01","message":"general error","info":null}
    }   
}

func myFunc() error {
    return typego.NewError("01", "general error")
}
```

```go
typego.NewError("01", "general error").SetHttpStatus(500).AddInfo("raw error 1", "raw error 2").AddInfo("raw error 3")
// output
// error: {"code":"01","message":"general error","info":["raw error 1","raw error 2","raw error 3"],"http_status":500}
```

You can also generate new `typego.Error` from an `error` string:

```go
err := errors.New("error: {\"code\":\"01\",\"message\":\"general error\",\"http_status\":500,\"info\":[\"raw info 1\",\"raw info 2\"],\"rpc_status\":13}")
typegoError := typego.NewErrorFromError(err)

fmt.Println(typegoError.GetCode()) // 01
fmt.Println(typegoError.GetMessage()) // general error
fmt.Println(typegoError.GetInfo()[1]) // raw info 2
```



## Release

### Changelog

Read at [CHANGELOG.md](https://github.com/dalikewara/typego/blob/master/CHANGELOG.md)

### Credits

Copyright &copy; 2023 [Dali Kewara](https://www.dalikewara.com)

### License

[MIT License](https://github.com/dalikewara/typego/blob/master/LICENSE)
