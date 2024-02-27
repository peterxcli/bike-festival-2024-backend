package main

import (
	"bikefest/pkg/bootstrap"
	"bikefest/pkg/model"
	"bikefest/pkg/service"
	"context"
	"log"
)

func main() {
	// init config
	app := bootstrap.App()

	// init services
	userService := service.NewUserService(app.Conn, app.Cache)
	notifyService := service.NewAsynqService(app.AsynqClient, app.AsynqInspector, app.Env)
	ctx := context.Background()
	usersWithEvents, err := userService.FindAllUserSubscribedEvents(ctx)
	if err != nil {
		panic(err)
	}
	for _, user := range usersWithEvents {
		for _, event := range user.Events {
			err := notifyService.EnqueueEventNotification(ctx, user.ID, *event.ID, event.EventTimeStart.Format(model.EventTimeLayout))
			if err != nil {
				log.Fatal("user: ", user.ID, " event: ", *event.ID, " error: ", err)
			}
		}
	}
}
