package main

import (
	"fmt"
	"log"
	"os"

	packer "github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2021-04-30/client/packer_service"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"

	"github.com/go-openapi/runtime/client"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2021-04-30/models"
	"github.com/hashicorp/hcp-sdk-go/httpclient"
)

// Parameters contains all the posisble inputs needed to make API requests.
type Parameters struct {
	// ID of the HCP Organization
	OrganizationID string
	// ID of the HCP Project
	ProjectID string
	// Human-readable Bucket name
	BucketSlug string
	// Human-readable Channel name
	ChannelSlug string
}

func main() {
	bucketSlug := os.Getenv("INPUT_BUCKET")
	channelSlug := os.Getenv("INPUT_CHANNEL")
	orgID := os.Getenv("HCP_ORGANIZATION_ID")
	projID := os.Getenv("HCP_PROJECT_ID")

	// Set params
	params := Parameters{
		OrganizationID: orgID,
		ProjectID:      projID,
		BucketSlug:     bucketSlug,
		ChannelSlug:    channelSlug,
	}

	// Initialize SDK client
	client, err := httpclient.New(httpclient.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Get latest Iteration in Bucket
	latestIterationID, err := getLatestIteration(client, &params)
	if err != nil {
		log.Fatal(err)
	}

	if channelExists(client, &params) {
		// Update Channel to latest Iteration
		err = updateChannel(client, &params, latestIterationID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Channel \"%s\" updated successfully!", channelSlug)
	} else {
		// Create Channel and set to latest Iteration
		createChannel(client, &params, latestIterationID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Channel \"%s\" created successfully!", channelSlug)
	}
}

// getLatestIteration takes returns the ID of the latest Iteration in the Bucket
func getLatestIteration(client *client.Runtime, params *Parameters) (string, error) {
	// Initialize packer client
	packerClient := packer.New(client, nil)

	// Initialize and set request params
	reqParams := packer.NewPackerServiceGetBucketParams()
	reqParams.LocationOrganizationID = params.OrganizationID
	reqParams.LocationProjectID = params.ProjectID
	reqParams.BucketSlug = params.BucketSlug

	// Send request
	resp, err := packerClient.PackerServiceGetBucket(reqParams, nil)
	if err != nil {
		return "", err
	}

	return resp.Payload.Bucket.LatestIteration.ID, nil
}

// channelExists returns true if the Channel exists
func channelExists(client *client.Runtime, params *Parameters) bool {
	// Initialize packer client
	packerClient := packer.New(client, nil)

	// Initialize and set request params
	reqParams := packer.NewPackerServiceGetChannelParams()
	reqParams.LocationOrganizationID = params.OrganizationID
	reqParams.LocationProjectID = params.ProjectID
	reqParams.BucketSlug = params.BucketSlug
	reqParams.Slug = params.ChannelSlug

	// Send request
	_, err := packerClient.PackerServiceGetChannel(reqParams, nil)
	return err == nil
}

// createChannel creates the specified Channel
func createChannel(client *client.Runtime, params *Parameters, iteration string) error {
	// Initialize packer client
	packerClient := packer.New(client, nil)

	// Initialize and set Location object
	var location cloud.HashicorpCloudLocationLocation
	location.OrganizationID = params.OrganizationID
	location.ProjectID = params.ProjectID

	// Initialize and set Body object
	var body models.HashicorpCloudPackerCreateChannelRequest
	body.Location = &location
	body.BucketSlug = params.BucketSlug
	body.Slug = params.ChannelSlug
	body.IterationID = iteration

	// Initialize and set request params
	reqParams := packer.NewPackerServiceCreateChannelParams()
	reqParams.LocationOrganizationID = params.OrganizationID
	reqParams.LocationProjectID = params.ProjectID
	reqParams.BucketSlug = params.BucketSlug
	reqParams.SetBody(&body)

	// Send request
	_, err := packerClient.PackerServiceCreateChannel(reqParams, nil)
	return err
}

// updateChannel points the Channel to the latest Iteration available
func updateChannel(client *client.Runtime, params *Parameters, iteration string) error {
	// Initialize packer client
	packerClient := packer.New(client, nil)

	// Initialize and set Location object
	var location cloud.HashicorpCloudLocationLocation
	location.OrganizationID = params.OrganizationID
	location.ProjectID = params.ProjectID

	// Initialize and set Body object
	var body models.HashicorpCloudPackerUpdateChannelRequest
	body.Location = &location
	body.BucketSlug = params.BucketSlug
	body.Slug = params.ChannelSlug
	body.IterationID = iteration

	// Initialize and set request params
	reqParams := packer.NewPackerServiceUpdateChannelParams()
	reqParams.LocationOrganizationID = params.OrganizationID
	reqParams.LocationProjectID = params.ProjectID
	reqParams.BucketSlug = params.BucketSlug
	reqParams.Slug = params.ChannelSlug
	reqParams.SetBody(&body)

	// Send request
	_, err := packerClient.PackerServiceUpdateChannel(reqParams, nil)
	return err
}
