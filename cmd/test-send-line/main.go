package main

import (
	"bikefest/pkg/bootstrap"
	"bikefest/pkg/service"
	"context"
	"fmt"
)

func main() {
	// init config
	app := bootstrap.App()

	// init services
	// userService := service.NewUserService(app.Conn, app.Cache)
	notifyService := service.NewAsynqService(app.AsynqClient, app.AsynqInspector, app.Env)
	ctx := context.Background()

	fmt.Println("Enter your user ID: ")
	var userID string
	_, err := fmt.Scan(&userID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter your event ID: ")
	var eventID string
	_, err = fmt.Scan(&eventID)
	if err != nil {
		panic(err)
	}
	println("Enter your event start time: ")
	var eventStartTime string
	_, err = fmt.Scan(&eventStartTime)
	if err != nil {
		panic(err)
	}
	err = notifyService.EnqueueEventNotification(ctx, userID, eventID, eventStartTime)
	if err != nil {
		panic(err)
	}
}
