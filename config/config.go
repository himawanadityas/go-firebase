package config

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Connection() (context.Context, *firebase.App) {
	option := option.WithCredentialsFile("go-firebase/firebaseKey.json")
	ctx := context.Background()
	conf := &firebase.Config{DatabaseURL:"https://data-118cf.firebaseio.com"}
	apps, err := firebase.NewApp(ctx, conf, option)
	if err != nil {
		fmt.Println(err)
	}
	return ctx, apps
}