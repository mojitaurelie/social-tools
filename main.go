package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"social-tools/config"
	"social-tools/data"
	"social-tools/scheduler"
	"social-tools/server"
)

func main() {
	fmt.Printf("social-tools v0.1.0 | Aurelie Delhaie | %s/%s\n\n", runtime.Version(), runtime.GOARCH)

	cpath := flag.String("config", "./config.yml", "The path to the configuration file")
	dbPath := flag.String("db", ":memory:", "The path to the SQLite storage")
	flag.Parse()

	err := config.Load(*cpath)
	if err != nil {
		log.Fatalf("CRITICAL %s", err)
	}
	err = data.Initialize(*dbPath)
	if err != nil {
		log.Fatalf("CRITICAL %s", err)
	}
	scheduler.Start()
	server.Serve()
}
