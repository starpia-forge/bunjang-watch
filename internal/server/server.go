package server

import (
	"io"
	"net/http"
	"os"
	"time"
)

const (
	DefaultServerAddress = "127.0.0.1:8080"
	DefaultIdleTimeout   = time.Minute
	DefaultReadTimeout   = time.Minute
	DefaultWriteTimeout  = time.Minute
)

type Server struct {
	Address      string
	Mode         string
	Logger       io.Writer
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Option func(*Server)

func WithServerAddress(address string) Option {
	return func(s *Server) {
		s.Address = address
	}
}

func WithServerMode(mode string) Option {
	return func(s *Server) {
		s.Mode = mode
	}
}

func WithServerLogger(logger io.Writer) Option {
	return func(s *Server) {
		s.Logger = logger
	}
}

func WithServerIdleTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.IdleTimeout = timeout
	}
}

func WithServerReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.ReadTimeout = timeout
	}
}

func WithServerWriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.WriteTimeout = timeout
	}
}

func NewServer(opts ...Option) (*http.Server, error) {
	s := &Server{
		Address:      DefaultServerAddress,
		IdleTimeout:  DefaultIdleTimeout,
		ReadTimeout:  DefaultReadTimeout,
		WriteTimeout: DefaultWriteTimeout,
		Logger:       os.Stdout,
	}

	for _, opt := range opts {
		opt(s)
	}

	return &http.Server{
		Addr:         s.Address,
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  s.IdleTimeout,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}, nil
}
