package main

import (
	"context"
	"fbchatbot/config"
	"fbchatbot/handler"
	"fbchatbot/libs/fb"
	"fbchatbot/logger"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go.uber.org/zap"
)

var configF = flag.String("config", "config.yml", "input config path, --config=config.yml")

func main() {
	flag.Parse()
	cfg := config.NewConfig(*configF)
	if cfg == nil {
		return
	}
	logger.InitLog(cfg.LogPath)
	ctx, cancel := context.WithCancel(context.Background())
	instance := fb.NewFacebookClient(cfg,
		handler.MessageReceived,
		handler.PostbackReceived,
		handler.DeliveryReceived)
	// set URL for your webhook and directly use msng as http Handler
	http.Handle("/mychatbot", instance.Client)
	listenAddr := fmt.Sprintf(":%d", cfg.Port)
	srv := http.Server{
		Addr:    listenAddr,
		Handler: http.DefaultServeMux,
	}
	var wg sync.WaitGroup
	wg.Add(1)
	SetupCloseHandler(cancel, &wg, &srv, ctx)

	// fb bot reply queue
	handler.MessageChan = make(chan handler.MsgTemplate, 1)
	wg.Add(1)
	go handler.SendFBMsg(ctx, &wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Logger.Info("bot service start",
			zap.String("listen", listenAddr))

		err := srv.ListenAndServe()
		if err != nil {
			logger.Logger.Error("bot service http stopped", zap.Error(err))
		}
	}()

	wg.Wait()
	logger.Logger.Info("bot service stop")
}

func SetupCloseHandler(cancel context.CancelFunc, wg *sync.WaitGroup, srv *http.Server, ctx context.Context) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		defer wg.Done()
		<-c
		cancel()
		err := srv.Shutdown(ctx)
		if err != nil {
			logger.Logger.Error("http stop error", zap.Error(err))
			return
		}
		logger.Logger.Info("\r- Ctrl+C pressed in Terminal")
	}()
}
