package handlers

import (
	"context"
	"regexp"

	"github.com/go-redis/redis/v8"
)

var (
	redirectURLre = regexp.MustCompile(`^\/([\w-]+)$`)
	slugRe        = regexp.MustCompile(`^[\w-]+$`)
	ctx           = context.TODO()
)

type handler struct {
	DB *redis.Client
}

func New(db *redis.Client) handler {
	return handler{db}
}
