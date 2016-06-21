package acl

import (
	"net/http"
	"sync"
)

type Policy interface {
	CheckAccessActor(actor, role string) error
	CheckAccessHTTP(req *http.Request, role string) error
}

var (
	once sync.Once
)
