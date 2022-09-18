Flow development

Domain:
- Entity
- Repository
- UseCase

Infrastructure : 
- database

Delivery :
- Container <br/>
Wrapping config,repository,usecase,shared,etc. You maybe wasted time for consider between interface
- Http Handler (server/http/handler) <br/>
Handler usecase for http request
- Middleware (optional) <br/>
For authentication header
- Router Handler
- Handler
- http
- Server
- main

How to run : <br>
go run cmd/app/main.go