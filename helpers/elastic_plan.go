package helpers

import "github.com/elastic/cloud-sdk-go/pkg/models"

// PrepareElasticSystemSettings accepts the needed variables and returns pointer to
// models.ElasticsearchSystemSettings
func PrepareElasticSystemSettings(autoCreateIndex bool, destructiveRequiresName bool, enableCloseIndex bool,
	useDiskThreshold bool) *models.ElasticsearchSystemSettings {
	// Initialize ElasticsearchSystemSettings
	esSettings := models.ElasticsearchSystemSettings{
		AutoCreateIndex:              &autoCreateIndex,         // Allow users to manage this setting
		DefaultShardsPerIndex:        1,                        // Using default
		DestructiveRequiresName:      &destructiveRequiresName, // Disables wildcard deletions. Should be true and users should control it
		EnableCloseIndex:             &enableCloseIndex,        // Should always be false. Refer to https://www.elastic.co/guide/en/cloud-enterprise/current/ece-add-user-settings.html#ece-change-user-settings-examples
		MonitoringCollectionInterval: 10,                       // defaults to 10 seconds
		MonitoringHistoryDuration:    "1d",                     // defaults to 7 days
		ReindexWhitelist:             nil,                      // TODO: check if we should allow this option during creation for whitelisting remote clusters
		Scripting:                    nil,                      // Not mandatory, use only when restricting scripting use is needed
		UseDiskThreshold:             &useDiskThreshold,        // Allow users to manage this setting
		// WatcherTriggerEngine:         "", // default to scheduler so not setting it through code
	}
	return &esSettings
}
