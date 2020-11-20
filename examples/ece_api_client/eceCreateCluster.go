package main

import (
	"encoding/json"
	"github.com/clutonik/elastic-cloud-tools/eceplans"
	"github.com/clutonik/elastic-cloud-tools/helpers"
	"github.com/clutonik/elastic-cloud-tools/utils"
	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi"
	"github.com/elastic/cloud-sdk-go/pkg/auth"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Evaluate Environment variables
	envVars := []string{"ECE_USER", "ECE_PASSWORD", "ECE_URL", "ECE_API_VERBOSE"}
	envMap, err := utils.LookupEnvVars(envVars)
	if err != nil {
		log.Fatalf("error while evaluating environment variables: %v", err)
	}

	// Create new auth config to use ECE credentials
	configuration := auth.Config{Username: envMap["ECE_USER"],
		Password: envMap["ECE_PASSWORD"],
	}

	// New AuthWriter
	eceAuthWriter, _ := auth.NewAuthWriter(configuration)

	// Parse Verbose flag from environment variable
	verboseEnabled, err := strconv.ParseBool(envMap["ECE_API_VERBOSE"])
	if err != nil {
		log.Printf("could not fetch verbose flag from environment variables: %v", err)
	}

	// Create new ECE API instance
	ece, err := api.NewAPI(api.Config{AuthWriter: eceAuthWriter,
		Client: new(http.Client),
		Host:   envMap["ECE_URL"],
		VerboseSettings: api.VerboseSettings{
			Device:     os.Stdout,
			Verbose:    verboseEnabled, // TODO: Write test case for this value
			RedactAuth: false,
		},
		ErrorDevice: os.Stdout,
	})
	if err != nil {
		log.Fatalf("could not create ECE API instance: %v", err)
	}

	var encoder = json.NewEncoder(os.Stdout)

	// Elasticsearch system settings
	autoCreateIndex := true
	destructiveRequiresName := true
	enableCloseIndex := false
	usediskThreshold := true

	// Prepare elasticsearch system settings
	esSettings := helpers.PrepareElasticSystemSettings(autoCreateIndex,
		destructiveRequiresName,
		enableCloseIndex,
		usediskThreshold)

	// Elasticsearch configuration
	esConfiguration := models.ElasticsearchConfiguration{
		Curation:                 nil,        // Only when using hot-warm templates
		DockerImage:              "",         // will not be exposed as controlling this will not be allowed
		EnabledBuiltInPlugins:    nil,        // plugins to add
		NodeAttributes:           nil,        // TODO: check the purpose of these attributes
		SystemSettings:           esSettings, //
		UserBundles:              nil,
		UserPlugins:              nil,
		UserSettingsJSON:         nil,
		UserSettingsOverrideJSON: nil,
		UserSettingsOverrideYaml: "",
		UserSettingsYaml:         "",
		Version:                  "",
	}

	// Elasticsearch cluster topology
	esTopology := models.ElasticsearchClusterTopologyElement{
		AllocatorFilter:         nil,
		Elasticsearch:           nil,
		InstanceConfigurationID: "",
		MemoryPerNode:           0,
		NodeConfiguration:       "",
		NodeCountPerZone:        0,
		NodeType:                nil,
		Size:                    nil,
		ZoneCount:               0,
	}

	// Elasticsearch cluster plan
	esPlan := models.ElasticsearchClusterPlan{
		ClusterTopology:    nil,
		DeploymentTemplate: nil,
		Elasticsearch:      nil,
		TiebreakerOverride: nil,
		TiebreakerTopology: nil,
		Transient:          nil,
		ZoneCount:          0,
	}

	// Elasticsearch payload
	payloads := []*models.ElasticsearchPayload{&models.ElasticsearchPayload{
		DisplayName: "test",
		Plan:        nil,
		RefID:       nil,
		Region:      nil,
		Settings:    nil,
	}}

	// Deployment Resources
	resources := models.DeploymentCreateResources{
		Apm:              nil,
		Appsearch:        nil,
		Elasticsearch:    nil,
		EnterpriseSearch: nil,
		Kibana:           nil,
	}

	// Deployment request
	req := models.DeploymentCreateRequest{
		Metadata:  nil, // metadata like system owned cluster
		Name:      "test",
		Resources: nil,
		Settings:  nil,
	}

	res, err := deploymentapi.Create(deploymentapi.CreateParams{
		API:       ece,
		Request:   nil,
		RequestID: "",
		Overrides: nil,
	})
	if err := encoder.Encode(res); err != nil {
		panic(err)
	}

}
