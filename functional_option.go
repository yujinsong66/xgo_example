package main

import (
	"fmt"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
}

func (server *Server) WithOptions(opts ...Option) *Server {
	for _, opt := range opts {
		opt.apply(server)
	}
	return server
}
func (server *Server) String() string {
	return fmt.Sprintf("Server: addr: %s, port: %d, protpcol: %s, timeout: %d, maxconn: %d",
		server.Addr, server.Port, server.Protocol, server.Timeout, server.MaxConns)
}

type Option interface {
	apply(*Server)
}

type optionFunc func(*Server)

func (f optionFunc) apply(server *Server) {
	f(server)
}

func Protocol(protocol string) Option {
	return optionFunc(func(server *Server) {
		server.Protocol = protocol
	})
}

func Timeout(timeout time.Duration) Option {
	return optionFunc(func(server *Server) {
		server.Timeout = timeout
	})
}

func MaxConn(maxConn int) Option {
	return optionFunc(func(server *Server) {
		server.MaxConns = maxConn
	})
}

func main() {
	server := &Server{Addr: "http://127.0.0.1", Port: 8080}
	opts := []Option{
		Protocol("http"),
		Timeout(100),
		MaxConn(1000),
	}
	server.WithOptions(opts...)
	fmt.Println(server)
}
