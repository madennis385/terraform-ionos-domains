package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"ionosdomains/gen"
)

func main() {
	client, err := ionosdomains.NewClientWithResponses(
		"https://api.ionos.com/domains", // Adjust if needed
		ionosdomains.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer YOUR_API_KEY")
			return nil
		}),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Adjust this to a valid method from the spec
	resp, err := client.ListDomainsWithResponse(context.Background())
	if err != nil {
		log.Fatalf("API call failed: %v", err)
	}

	fmt.Printf("Domains: %+v\n", resp.JSON200)
}

