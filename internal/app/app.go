package app

import (
	"context"
	"errors"
	"github.com/starpia-forge/bunjang-watch/internal/server"
	"github.com/starpia-forge/bunjang-watch/internal/watcher"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	server  *http.Server
	manager *watcher.WatcherManager
}

func New() (*App, error) {
	srv, err := server.NewServer()
	if err != nil {
		return nil, err
	}

	return &App{
		server:  srv,
		manager: watcher.NewWatcherManager(),
	}, nil
}

func (a *App) Run() error {
	appCtx, appCancel := context.WithCancel(context.Background())
	defer appCancel()

	// 서버 실행
	go func() {
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// 치명적인 에러일 경우 로그 출력 후 종료
			// log.Printf("ListenAndServe error: %v", err)
			appCancel()
		}
	}()

	// 시스템 시그널 처리 (Ctrl+C 등)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case <-shutdown:
		appCancel()
	case <-appCtx.Done():
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := a.server.Shutdown(shutdownCtx); err != nil {
		return err
	}

	return nil
}
