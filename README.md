# web-template
>Template for rapid construction of Web services based on gin framework

# Introduction
Gin has been able to help us quickly implement some web functions, but I think if I want to build a complete web project quickly, it still needs some time and more steps.

So I encapsulated it again, so that I could use it more succinctly and build faster.

**For example, I automatically map the request parameters and corresponding parameters with the structure, and you only need to focus on what parameters you need and what parameters to respond to.**

And you don't have to worry about being difficult to use, because you can see the use style of Gin in many places, which is only easier to use than Gin
# How to use
You just need to select the following code in the template:
```
.
│ 
│ main.go 
└─routes
│   └─routes.go
├─handlers  
│   └─hello.go
└─proto
│   └─hello.go
└─config
    ├─application.json
    └─config.go

```
This example creates a web server in a custom configuration, default load config path: %project%/config/application.json, 
However, you can also specify the path of the configuration file by passing parameters through the command function, like: 
```
go run main.go -c app.json
```

And you can choose to create a default web server with port 8080 and environment prod, like:
```go
func main() {
    server.DefaultServer().Router(routes.Routes).Run() 
}
```

It is suggested that you can explore the functions of web template according to the usage habits of Gin