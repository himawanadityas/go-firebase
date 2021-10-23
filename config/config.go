package config

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Connection(){
	option := option.WithCredentialFiles("firebaseKey.json")
	ctx := context.Background()
	conf := &firebase.config{DatabaseURL:"https://data-118cf.firebaseio.com"}
	apps, err := firebase.NewApp(ctx, conf, option)
	if err != nil {
		fmt.Println(err)
	}
	initDB, err := apps.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db := initDB.NewRef("my-apps")
	return db, apps
}