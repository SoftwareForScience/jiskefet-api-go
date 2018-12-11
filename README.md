# Jiskefet Go Client API
This is a Go client API for Jiskefet based on the go-swagger code generator (https://goswagger.io)


## Setup
```
go get -u -v github.com/PascalBoeschoten/jiskefet-api-go
```


## Running the example code
If your api is at `http://myhost.server.address/api`

```
cd $GOPATH/src/github.com/PascalBoeschoten/jiskefet-api-go
export JISKEFET_HOST=myhost.server.address
export JISKEFET_PATH=api
export JISKEFET_API_TOKEN=jnk5vh43785ycj4gdvlvm84fg...
go run main.go
```


## Code generation
To (re)generate the client code:
```
$ ~/go/bin/swagger generate client --spec=swagger.yaml
```
