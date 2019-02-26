package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-openapi/strfmt"
	//"github.com/go-openapi/spec"

	"github.com/SoftwareForScience/jiskefet-api-go/client/runs"
	runsclient "github.com/SoftwareForScience/jiskefet-api-go/client/runs"
	"github.com/SoftwareForScience/jiskefet-api-go/models"

	//apiclient "github.com/SoftwareForScience/jiskefet-api-go/client"
	httptransport "github.com/go-openapi/runtime/client"
)

func main() {
	fmt.Println("Hello Jiskefet")

	hostUrl := os.Getenv("JISKEFET_HOST")
	apiPath := os.Getenv("JISKEFET_PATH")
	apiToken := os.Getenv("JISKEFET_API_TOKEN")
	fmt.Printf("JISKEFET_HOST:      %s\n", hostUrl)
	fmt.Printf("JISKEFET_PATH:  %s\n", apiPath)
	fmt.Printf("JISKEFET_API_TOKEN: %s\n", apiToken)

	// create the API client
	//client := apiclient.New(httptransport.New("", "", nil), strfmt.Default)
	client := runsclient.New(httptransport.New(hostUrl, apiPath, nil), strfmt.Default)

	// make the authenticated request to get all items
	bearerTokenAuth := httptransport.BearerToken(apiToken)
	// basicAuth := httptransport.BasicAuth(os.Getenv("API_USER"), os.Getenv("API_PASSWORD"))
	// apiKeyQueryAuth := httptransport.APIKeyAuth("apiKey", "query", os.Getenv("API_KEY"))
	// apiKeyHeaderAuth := httptransport.APIKeyAuth("X-API-TOKEN", "header", os.Getenv("API_KEY"))

	{
		startt, err := time.Parse(time.RFC3339, "2007-03-01T13:40:45Z")
		if err != nil {
			fmt.Println(err)
		}

		endt, err := time.Parse(time.RFC3339, "2016-04-01T10:22:02Z")
		if err != nil {
			fmt.Println(err)
		}

		start := strfmt.DateTime(startt)
		end := strfmt.DateTime(endt)
		runType := "my-run-type"
		runQuality := "my-run-quality"
		activityId := "go-api"
		nDetectors := int64(123)
		nFlps := int64(200)
		nEpns := int64(1000)
		nTimeframes := int64(1231)
		nSubtimeframes := int64(12312)
		bytesReadOut := int64(1024 * 1024)
		bytesTimeframeBuilder := int64(512 * 1024)

		params := runs.NewPostRunsParams()
		params.CreateRunDto = new(models.CreateRunDto)
		params.CreateRunDto.TimeO2Start = &start
		params.CreateRunDto.TimeTrgStart = &start
		params.CreateRunDto.TimeO2End = &end
		params.CreateRunDto.TimeTrgEnd = &end
		params.CreateRunDto.RunType = &runType
		params.CreateRunDto.RunQuality = &runQuality
		params.CreateRunDto.ActivityID = &activityId
		params.CreateRunDto.NDetectors = &nDetectors
		params.CreateRunDto.NFlps = &nFlps
		params.CreateRunDto.NEpns = &nEpns
		params.CreateRunDto.NTimeframes = &nTimeframes
		params.CreateRunDto.NSubtimeframes = &nSubtimeframes
		params.CreateRunDto.BytesReadOut = &bytesReadOut
		params.CreateRunDto.BytesTimeframeBuilder = &bytesTimeframeBuilder

		_, err = client.PostRuns(params, bearerTokenAuth)
		if err != nil {
			fmt.Println(err)
		}
	}

	{
		params := runs.NewGetRunsParams()
		resp, err := client.GetRuns(params, bearerTokenAuth)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v\n", resp.Payload)
	}
}
