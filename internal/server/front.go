package server

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed static/*
var content embed.FS

func SPAHandler() http.Handler {
	return http.FileServer(http.FS(StaticFS()))
}

func StaticFS() fs.FS {
	sub, err := fs.Sub(content, "static")
	if err != nil {
		panic(err)
	}
	return sub
}
