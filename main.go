package main

import (
	"fmt"
	"log"
	"math/rand"
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

	hostURL := os.Getenv("JISKEFET_HOST")
	apiPath := os.Getenv("JISKEFET_PATH")
	apiToken := os.Getenv("JISKEFET_API_TOKEN")
	fmt.Printf("JISKEFET_HOST: %s\n", hostURL)
	fmt.Printf("JISKEFET_PATH: %s\n", apiPath)
	fmt.Printf("JISKEFET_API_TOKEN: %s\n", apiToken)

	// create the API client
	//client := apiclient.New(httptransport.New("", "", nil), strfmt.Default)
	client := runsclient.New(httptransport.New(hostURL, apiPath, nil), strfmt.Default)

	// make the authenticated request to get all items
	bearerTokenAuth := httptransport.BearerToken(apiToken)
	// basicAuth := httptransport.BasicAuth(os.Getenv("API_USER"), os.Getenv("API_PASSWORD"))
	// apiKeyQueryAuth := httptransport.APIKeyAuth("apiKey", "query", os.Getenv("API_KEY"))
	// apiKeyHeaderAuth := httptransport.APIKeyAuth("X-API-TOKEN", "header", os.Getenv("API_KEY"))

	{
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		runNumber := int64(r1.Intn(10000))

		// Start run
		{
			startt, err := time.Parse(time.RFC3339, "2007-03-01T13:40:45Z")
			if err != nil {
				fmt.Println(err)
			}
			start := strfmt.DateTime(startt)
			runType := "TECHNICAL"
			activityID := "go-api-example"
			nDetectors := int64(123)
			nFlps := int64(200)
			nEpns := int64(1000)

			params := runs.NewPostRunsParams()
			params.CreateRunDto = new(models.CreateRunDto)
			params.CreateRunDto.O2StartTime = &start
			params.CreateRunDto.TrgStartTime = &start
			params.CreateRunDto.ActivityID = &activityID
			params.CreateRunDto.NDetectors = &nDetectors
			params.CreateRunDto.NEpns = &nEpns
			params.CreateRunDto.NFlps = &nFlps
			params.CreateRunDto.RunNumber = &runNumber
			params.CreateRunDto.RunType = &runType

			_, err = client.PostRuns(params, bearerTokenAuth)
			if err != nil {
				fmt.Println(err)
			}
		}

		// End run
		{
			endt, err := time.Parse(time.RFC3339, "2016-04-01T10:22:02Z")
			if err != nil {
				fmt.Println(err)
			}
			end := strfmt.DateTime(endt)
			runQuality := "my-run-quality"

			params := runs.NewPatchRunsIDParams()
			params.PatchRunDto = new(models.PatchRunDto)
			params.PatchRunDto.O2EndTime = &end
			params.PatchRunDto.TrgEndTime = &end
			params.PatchRunDto.RunQuality = &runQuality
			params.SetID(runNumber)

			_, err = client.PatchRunsID(params, bearerTokenAuth)
			if err != nil {
				fmt.Println(err)
			}
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
