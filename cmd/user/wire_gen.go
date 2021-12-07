// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/stonecutter/blog-microservices/api/protobuf"
	"github.com/stonecutter/blog-microservices/internal/pkg/config"
	"github.com/stonecutter/blog-microservices/internal/pkg/dbcontext"
	"github.com/stonecutter/blog-microservices/internal/user"
)

// Injectors from wire.go:

func InitServer(conf *config.Config) (protobuf.UserServiceServer, error) {
	db, err := dbcontext.NewUserDB(conf)
	if err != nil {
		return nil, err
	}
	repository := user.NewRepository(db)
	userServiceServer := user.NewServer(repository)
	return userServiceServer, nil
}