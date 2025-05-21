package util

import (
	"context"
	"io"
	"net/http"
)

// PostJSON sends a POST request with JSON payload to the given URL.
func PostJSON(ctx context.Context, client *http.Client, url string, body io.Reader, customHeader []map[string]string) (*http.Response, error) {
	return post(ctx, client, url, "application/json", body, customHeader)
}

func post(ctx context.Context, client *http.Client, url, bodyType string, body io.Reader, customHeader []map[string]string) (*http.Response, error) {
	return request(ctx, client, http.MethodPost, url, bodyType, body, customHeader)
}

func request(ctx context.Context, client *http.Client, method, url, bodyType string, body io.Reader, customHeader []map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if bodyType != "" {
		req.Header.Set("Content-Type", bodyType)
	}
	if len(customHeader) > 0 {
		for _, h := range customHeader {
			for k := range h {
				req.Header.Set(k, h[k])
			}
		}
	}
	return client.Do(req.WithContext(ctx))
}

// Drain consumes and closes the response's body to make sure that the
// HTTP client can reuse existing connections.
func Drain(r *http.Response) {
	io.Copy(io.Discard, r.Body)
	r.Body.Close()
}

func ReadAll(r io.Reader) string {
	if r == nil {
		return ""
	}
	bs, err := io.ReadAll(r)
	if err != nil {
		return ""
	}
	return string(bs)
}
