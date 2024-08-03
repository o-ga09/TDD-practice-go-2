package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/o-ga09/note-app-backendapi/api"
	"github.com/o-ga09/note-app-backendapi/db/dao"
	"github.com/o-ga09/note-app-backendapi/db/db"
	"github.com/o-ga09/note-app-backendapi/handler"
	"github.com/o-ga09/note-app-backendapi/pkg/logger"
	"github.com/o-ga09/note-app-backendapi/pkg/middleware"
	"github.com/o-ga09/note-app-backendapi/services/note"
	"github.com/o-ga09/note-app-backendapi/services/user"
)

func main() {
	ctx := context.Background()
	logger.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	db := db.Connect(ctx)
	if db == nil {
		panic(err)
	}

	noteRepo := dao.NewNoteDao(db)
	noteService := note.NewNoteService(noteRepo)

	userRepo := dao.NewUserDao(db)
	userService := user.NewUserService(userRepo)

	h := handler.NewHandler(*noteService, *userService)
	handler, err := api.NewServer(h,
		api.WithMiddleware(middleware.AddID()),
		api.WithMiddleware(middleware.WithTimeout()),
		api.WithMiddleware(middleware.RequestLogger()),
	)

	slog.Info("starting server")
	go func() {
		err = http.Serve(listen, handler)
		if err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("stopping srever")
}
