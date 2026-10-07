package monitor

import (
	"net/http"
	"time"

	"github.com/Rugved-dev18/gowatch/internal/result"
)

func CheckURL(url string) result.Result {
	start := time.Now()

	response, err := http.Get(url)

	latency := time.Since(start)

	if err != nil {
		return result.Result{
			URL:     url,
			Latency: latency,
			Error:   err,
		}
	}

	defer response.Body.Close()

	return result.Result{
		URL:        url,
		StatusCode: response.StatusCode,
		Latency:    latency,
	}
}
