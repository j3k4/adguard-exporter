package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/j3k4/adguard-exporter/config"
	"github.com/j3k4/adguard-exporter/internal/adguard"
	"github.com/j3k4/adguard-exporter/internal/metrics"
	"github.com/j3k4/adguard-exporter/internal/server"
)

const (
	name = "adguard-exporter"
)

var (
	s *server.Server
)

func main() {
	conf := config.Load()

	metrics.Init()

	initAdguardClient(conf.AdguardProtocol, conf.AdguardHostname, conf.AdguardUsername, conf.AdguardPassword, conf.AdguardPort, conf.Interval, conf.LogLimit, conf.RDnsEnabled, conf.InsecureTLSSkipVerify)
	initHttpServer(conf.ServerPort)

	handleExitSignal()
}

func initAdguardClient(protocol, hostname, username, password, port string, interval time.Duration, logLimit string, rdnsenabled bool, insecuretls bool) {
	client := adguard.NewClient(protocol, hostname, username, password, port, interval, logLimit, rdnsenabled, insecuretls)
	go client.Scrape()
}

func initHttpServer(port string) {
	s = server.NewServer(port)
	go s.ListenAndServe()
}

func handleExitSignal() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	s.Stop()
	fmt.Println(fmt.Sprintf("\n%s HTTP server stopped", name))
}
