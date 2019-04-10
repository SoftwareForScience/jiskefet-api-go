# Jiskefet Go Client API
This is a Go client API for Jiskefet based on the go-swagger code generator (https://goswagger.io)


## Installation
```
go get -u -v github.com/SoftwareForScience/jiskefet-api-go
```


## Setup for development
```
export GOPATH=/path/to/my/go/workspace
go get -u -v github.com/SoftwareForScience/jiskefet-api-go
cd $GOPATH/src/github.com/SoftwareForScience/jiskefet-api-go
```


## Running the example code
If the API is running at `http://myhost.server.address/api`
```
export GOPATH=/path/to/my/go/workspace
go get -u -v github.com/SoftwareForScience/jiskefet-api-go
cd $GOPATH/src/github.com/SoftwareForScience/jiskefet-api-go
export JISKEFET_HOST=myhost.server.address
export JISKEFET_PATH=api
export JISKEFET_API_TOKEN=jnk5vh43785ycj4gdvlvm84fg...
go run main.go
```


## Code generation
To (re)generate the client code:
```
go get -u -v github.com/go-swagger/go-swagger/cmd/swagger
cd $GOPATH/src/github.com/SoftwareForScience/jiskefet-api-go
~/go/bin/swagger generate client --spec=jiskefet-openapi-spec/openapi-spec.yaml
```
