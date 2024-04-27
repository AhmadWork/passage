package main

import (
	"net/http/httputil"
	"net/url"
	"sync"
)


type Backend struct {
  URL          *url.URL
  Alive        bool
  mux          sync.RWMutex
  Name string
  ReverseProxy *httputil.ReverseProxy
}

type ServerPool struct {
  backends map[string]*Backend
  current  uint64
}


