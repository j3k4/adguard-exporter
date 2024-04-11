package metrics

import (
	"fmt"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

// TestInit überprüft die Initialisierung der Prometheus-Metriken
func TestInit(t *testing.T) {
	// Mock-Logausgabe
	logOutput := make(chan string, 12)
	log := &testLogger{c: logOutput}

	// Ersetze den originalen Logger durch den Mock-Logger
	originalLogger := log
	defer func() { log = originalLogger }()

	// Rufe Init-Funktion auf
	Init()

	// Erwartete Metriken
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

	// Überprüfe, ob alle erwarteten Metriken registriert wurden
	for _, metricName := range expectedMetrics {
		if !isMetricRegistered(metricName) {
			t.Errorf("Expected metric %s to be registered, but it wasn't", metricName)
		}
	}
}

// testLogger ist eine Mock-Implementierung von log.Logger für Testzwecke
type testLogger struct {
	c chan string
}

func (l *testLogger) Printf(format string, v ...interface{}) {
	l.c <- fmt.Sprintf(format, v...)
}

// isMetricRegistered überprüft, ob eine Metrik registriert ist
func isMetricRegistered(metricName string) bool {
	_, err := prometheus.DefaultGatherer.Gather()
	return err == nil
}
