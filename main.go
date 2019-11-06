package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()

	conf := oauth2.Config{
		ClientID: "rpxapps",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://auth.rpxapps.com/auth/realms/rpx/protocol/openid-connect/auth",
			TokenURL: "https://auth.rpxapps.com/auth/realms/rpx/protocol/openid-connect/token",
		},
		RedirectURL: "http://localhost:8080/hello",
	}

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	client.Get("...")

}
