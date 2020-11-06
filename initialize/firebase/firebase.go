package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/birnamwood/go-nuxt/config"
	"google.golang.org/api/option"
)

var fb *firebase.App
var message *messaging.Client

//Init firebase Initialize
func Init() {
	c := config.GetConfig()
	ctx := context.Background()
	opt := option.WithCredentialsFile(c.GetString("firebase.credential"))
	config := &firebase.Config{ProjectID: "golang-messaging"}
	var err error
	fb, err = firebase.NewApp(ctx, config, opt)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	message, err = fb.Messaging(ctx)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}

//GetFCM return Firebase Clound Messaging
func GetFCM() *messaging.Client {
	return message
}
