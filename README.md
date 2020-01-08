# Jiskefet Go Client API
This is a Go client API for Jiskefet based on the go-swagger code generator (https://goswagger.io)

Note: this repo contains a submodule `jiskefet-openapi-spec` that should point towards the spec that was used to generate the code.

## Setup
```
export GOPATH=/my/working/dir/
go get -u -v github.com/SoftwareForScience/jiskefet-api-go
```


## Running the example code
If your api is at `http://myhost.server.address/api`

```
cd $GOPATH/src/github.com/SoftwareForScience/jiskefet-api-go
export JISKEFET_HOST=myhost.server.address
export JISKEFET_PATH=api
export JISKEFET_API_TOKEN=jnk5vh43785ycj4gdvlvm84fg...
go run main.go
```


## Code generation
First, the yaml file needs to be generated from the api-doc, run the following command:
```
bash generate-yaml.sh url/of/api/doc
```

Note: sometimes the generation f^cks up, in this case, you'll nee to manually change the lines using editor.swagger.io.  


To (re)generate the client code:
```
go get -u -v github.com/go-swagger/go-swagger/cmd/swagger
cd $GOPATH/src/github.com/SoftwareForScience/jiskefet-api-go
git submodule init
git submodule update
go get -u -f ./..
$GOPATH/bin/swagger generate client --spec=doc.yaml
```
