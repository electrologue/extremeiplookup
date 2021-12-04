package extremeiplookup

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) (*Client, *http.ServeMux) {
	t.Helper()

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	client := NewClient("secret")
	client.baseURL, _ = url.Parse(server.URL)
	client.HTTPClient = server.Client()

	return client, mux
}

func testHandler(filename string) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(rw, fmt.Sprintf("unsupported method: %s", req.Method), http.StatusMethodNotAllowed)
			return
		}

		file, err := os.Open(filepath.Join("fixtures", filename))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		defer func() { _ = file.Close() }()

		_, err = io.Copy(rw, file)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func TestClient_Lookup_error(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/json/109.236.91.3", testHandler("fail.json"))

	ipInfo, err := client.Lookup(context.Background(), "109.236.91.3")
	require.Error(t, err)
	require.Nil(t, ipInfo)
}

func TestClient_Lookup(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/json/109.236.91.3", testHandler("success.json"))

	ipInfo, err := client.Lookup(context.Background(), "109.236.91.3")
	require.NoError(t, err)

	expected := &IPInfo{
		Query:        "109.236.91.3",
		IPType:       "Residential",
		Continent:    "Europe",
		CountryCode:  "NL",
		Country:      "Netherlands",
		Region:       "Zuid-Holland",
		City:         "Naaldwijk",
		Latitude:     "51.99417",
		Longitude:    "4.20972",
		IPName:       "customer.worldstream.nl",
		Organization: "WorldStream B.V.",
		ISP:          "WorldStream B.V.",
		Timezone:     "Europe/Amsterdam",
		UTCOffset:    "+01:00",
		Status:       "success",
	}

	assert.Equal(t, expected, ipInfo)
}
