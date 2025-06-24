package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func makeRequest(etag string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/data", nil)

	if etag != "" {
		req.Header.Set("If-None-Match", etag)
		fmt.Println("🔁 Sending If-None-Match:", etag)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("📥 Status:", resp.Status)

	if resp.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("📦 Response Body:", string(body))
	}

	if et := resp.Header.Get("ETag"); et != "" {
		fmt.Println("🧾 Received ETag:", et)
	}
}

func main() {
	fmt.Println("🔍 First request (fresh)")
	makeRequest("") // No ETag → server returns full content

	time.Sleep(6 * time.Second) // simulate second request later

	fmt.Println("\n🔍 Second request (with ETag)")
	makeRequest(`"v1"`) // simulate cached version
}
