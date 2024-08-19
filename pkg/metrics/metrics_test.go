package metrics_test

import "testing"

func TestFetchMetrics(t *testing.T) {
	t.Run("should return metrics data from Kubernetes API", func(t *testing.T) {
		got := FetchMetrics()
		want := "expected_metrics_data"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
