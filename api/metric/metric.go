package metric

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsApi(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}
