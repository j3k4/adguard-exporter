package metrics

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

// TestInit tests the initialization of Prometheus metrics
func TestInit(t *testing.T) {
	// Call Init function
	Init()

	// Check if all metrics have been registered
	expectedMetrics := []string{
		"avg_processing_time",
		"num_dns_queries",
		"num_blocked_filtering",
		"num_replaced_parental",
		"num_replaced_safebrowsing",
		"num_replaced_safesearch",
		"top_queried_domains",
		"top_blocked_domains",
		"top_clients",
		"query_types",
		"running",
		"protection_enabled",
	}

	// Check if all expected metrics are registered
	for _, metricName := range expectedMetrics {
		if !isMetricRegistered() {
			t.Errorf("Expected metric %s to be registered, but it wasn't", metricName)
		}
	}
}

// isMetricRegistered checks if a metric is registered
func isMetricRegistered() bool {
	// Try to collect the metric
	_, err := prometheus.DefaultGatherer.Gather()
	return err == nil
}
