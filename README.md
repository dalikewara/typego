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

## Usage

### Error

`typego.Error` compatible with `error` interface, so you can use it as an `error` handler. `typego.Error`
has several methods that can be used to construct error information:

```go
type Error interface {
    ChangeCode(code string) Error
    ChangeMessage(message string) Error
    AddInfo(info ...any) Error
    AddDebug(debug ...any) Error
    SetProcessID(processID string) Error
    SetProcessName(processName string) Error
    SetHttpStatus(httpStatus int) Error
    SetRPCStatus(rpcStatus int) Error
    GetProcessID() string
    GetProcessName() string
    GetCode() string
    GetMessage() string
    GetInfo() []string
    GetDebug() []string
    GetHttpStatus() int
    GetRPCStatus() int
    Log() Error
    Error() string
}
```

and it will generate the error information based on this structure:

```go
type errorModel struct {
    Level       string   `json:"level"`
    ProcessID   string   `json:"process_id,omitempty"`
    ProcessName string   `json:"process_name,omitempty"`
    Code        string   `json:"code"`
    Message     string   `json:"message"`
    Info        []string `json:"info"`
    HttpStatus  int      `json:"http_status,omitempty"`
    RPCStatus   int      `json:"rpc_status,omitempty"`
    Debug       []string `json:"debug,omitempty"`
}
```

For example:

```go
func main() {
    if err := myFunc(); err != nil {
        fmt.Println(err)
		
        // output
        // {"level":"error","code":"01","message":"general error","info":null}
    }   
}

func myFunc() error {
    return typego.NewError("01", "general error")
}
```

```go
typego.NewError("01", "general error").SetHttpStatus(500).AddInfo("raw error 1", "raw error 2").AddInfo("raw error 3")

// output
// {"level":"error","code":"01","message":"general error","info":["raw error 1","raw error 2","raw error 3"],"http_status":500}
```

You can log the error information by using `Log()` method:

```go
typego.NewError("01", "general error").Log()

// output
// {"level":"error","code":"01","message":"general error","info":null}
```

You can also generate new `typego.Error` from an `error`:

```go
err := errors.New("{\"code\":\"01\",\"message\":\"general error\",\"http_status\":500,\"info\":[\"raw info 1\",\"raw info 2\"],\"rpc_status\":13}")
typegoError := typego.NewErrorFromError(err)

fmt.Println(typegoError.GetCode()) // 01
fmt.Println(typegoError.GetMessage()) // general error
fmt.Println(typegoError.GetInfo()[1]) // raw info 2
```

> The `error.Error()` must have the same string format as `typego.Error.Error()`, otherwise, `typego.Error` will return incorrect value

#### Custom Error Log

You can overwrite the default error log handler by using `typego.SetCustomErrorLog(handler ErrorLogHandler)` function:

> The default error log handler is just a simple task to print the information to the std out

```go
errGeneral := typego.NewError("01", "general error")

errGeneral.Log()

// output
// {"level":"error","code":"01","message":"general error","info":null}

typego.SetCustomErrorLog(func(err typego.Error) {
    fmt.Println(fmt.Sprintf("hello i am a custom log! -> %+v", err))
	
    // or do something special here...
    // for example: send the log info to the Slack Channel, Kafka, etc
})

errGeneral.Log()

// output
// hello i am a custom log! -> {"level":"error","code":"01","message":"general error","info":null}
```

So, you can change the behavior of the logging as you want.

## Release

### Changelog

Read at [CHANGELOG.md](https://github.com/dalikewara/typego/blob/master/CHANGELOG.md)

### Credits

Copyright &copy; 2023 [Dali Kewara](https://www.dalikewara.com)

### License

[MIT License](https://github.com/dalikewara/typego/blob/master/LICENSE)
