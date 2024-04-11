package metrics

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

// TestInit testet die Initialisierung der Metriken
func TestInit(t *testing.T) {
	// Testen, ob alle Metriken registriert werden
	Init()

	// Überprüfen, ob jede Metrik registriert ist
	expectedMetrics := []struct {
		name  string
		metric *prometheus.GaugeVec
	}{
		{"avg_processing_time", AvgProcessingTime},
		{"num_dns_queries", DnsQueries},
		{"num_blocked_filtering", BlockedFiltering},
		{"num_replaced_parental", ParentalFiltering},
		{"num_replaced_safebrowsing", SafeBrowsingFiltering},
		{"num_replaced_safesearch", SafeSearchFiltering},
		{"top_queried_domains", TopQueries},
		{"top_blocked_domains", TopBlocked},
		{"top_clients", TopClients},
		{"query_types", QueryTypes},
		{"running", Running},
		{"protection_enabled", ProtectionEnabled},
	}

	for _, em := range expectedMetrics {
		if !contains(prometheus.DefaultRegisterer.(*prometheus.Registry).Collectors(), em.metric) {
			t.Errorf("expected metric %q to be registered", em.name)
		}
	}
}

// Hilfsfunktion, um zu überprüfen, ob ein Metric im Registry vorhanden ist
func contains(collectors []prometheus.Collector, needle prometheus.Collector) bool {
	for _, c := range collectors {
		if c == needle {
			return true
		}
	}
	return false
}
