package config

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Connection() (context.Context, *firebase.App) {
	option := option.WithCredentialsFile("firebaseKey.json")
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "data-118cf"}
	apps, err := firebase.NewApp(ctx, conf, option)
	if err != nil {
		fmt.Println(err)
	}
	return ctx, apps
}