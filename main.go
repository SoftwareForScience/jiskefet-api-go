package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-openapi/strfmt"
	//"github.com/go-openapi/spec"

	runsclient "github.com/PascalBoeschoten/jiskefet-api-go/client/runs"
	//apiclient "github.com/PascalBoeschoten/jiskefet-api-go/client"
	httptransport "github.com/go-openapi/runtime/client"
)

func main() {
	fmt.Println("Hello Jiskefet")

	// create the API client
	//client := apiclient.New(httptransport.New("", "", nil), strfmt.Default)
	client := runsclient.New(httptransport.New(os.Getenv("JISKEFET_HOST"), "", nil), strfmt.Default)

	// make the authenticated request to get all items
	bearerTokenAuth := httptransport.BearerToken(os.Getenv("JISKEFET_API_TOKEN"))
	// basicAuth := httptransport.BasicAuth(os.Getenv("API_USER"), os.Getenv("API_PASSWORD"))
	// apiKeyQueryAuth := httptransport.APIKeyAuth("apiKey", "query", os.Getenv("API_KEY"))
	// apiKeyHeaderAuth := httptransport.APIKeyAuth("X-API-TOKEN", "header", os.Getenv("API_KEY"))
	//resp, err := client.Operations.All(operations.AllParams{}, bearerTokenAuth)
	resp, err := client.GetRuns(nil, bearerTokenAuth)
	// resp, err := client.Operations.All(operations.AllParams{}, basicAuth)
	// resp, err := client.Operations.All(operations.AllParams{}, apiKeyQueryAuth)
	// resp, err := client.Operations.All(operations.AllParams{}, apiKeyHeaderAuth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}
