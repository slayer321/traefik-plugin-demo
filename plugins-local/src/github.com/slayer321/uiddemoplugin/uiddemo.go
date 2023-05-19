package uiddemoplugin

import (
	"context"
	"net/http"
	"fmt"

	"github.com/google/uuid"
	
)

const defaultHeader = "X-Traefik-Uuid"

type Config struct{
	HeaderName string
}

func CreateConfig() *Config {
	return &Config{
		HeaderName: defaultHeader,
	}
}

type UIDDemo struct {
	next http.Handler
	name string
	HeaderName string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.HeaderName) == 0{
		return nil, fmt.Errorf("HeaderName is Empty")
	}
	
	return &UIDDemo{
		next: next,
		name: name,
		HeaderName: config.HeaderName,
	}, nil
}

func (u *UIDDemo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	//rw.Write([]byte("Hello World"))

	id := uuid.New()
	req.Header.Set(u.HeaderName, id.String())

	rw.Header().Set(u.HeaderName, id.String())
	u.next.ServeHTTP(rw, req)
}
