package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-openapi/strfmt"
	//"github.com/go-openapi/spec"

	"github.com/PascalBoeschoten/jiskefet-api-go/client/runs"
	runsclient "github.com/PascalBoeschoten/jiskefet-api-go/client/runs"

	//apiclient "github.com/PascalBoeschoten/jiskefet-api-go/client"
	httptransport "github.com/go-openapi/runtime/client"
)

func main() {
	fmt.Println("Hello Jiskefet")

	hostUrl := os.Getenv("JISKEFET_HOST")
	apiPath := os.Getenv("JISKEFET_PATH")
	apiToken := os.Getenv("JISKEFET_API_TOKEN")
	fmt.Printf("JISKEFET_HOST:      %s\n", hostUrl)
	fmt.Printf("JISKEFET_API_PATH:  %s\n", apiPath)
	fmt.Printf("JISKEFET_API_TOKEN: %s\n", apiToken)

	// create the API client
	//client := apiclient.New(httptransport.New("", "", nil), strfmt.Default)
	client := runsclient.New(httptransport.New(hostUrl, apiPath, nil), strfmt.Default)

	// make the authenticated request to get all items
	bearerTokenAuth := httptransport.BearerToken(apiToken)
	// basicAuth := httptransport.BasicAuth(os.Getenv("API_USER"), os.Getenv("API_PASSWORD"))
	// apiKeyQueryAuth := httptransport.APIKeyAuth("apiKey", "query", os.Getenv("API_KEY"))
	// apiKeyHeaderAuth := httptransport.APIKeyAuth("X-API-TOKEN", "header", os.Getenv("API_KEY"))

	params := runs.NewGetRunsParams()
	resp, err := client.GetRuns(params, bearerTokenAuth)
	// resp, err := client.Operations.All(operations.AllParams{}, basicAuth)
	// resp, err := client.Operations.All(operations.AllParams{}, apiKeyQueryAuth)
	// resp, err := client.Operations.All(operations.AllParams{}, apiKeyHeaderAuth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}
