package main

import (
	"flag"
	"fmt"
	"log"
	"social-tools/config"
	"social-tools/scheduler"
	"social-tools/server"
)

func main() {
	fmt.Printf("social-tools v1.0.0\n\n")

	cpath := flag.String("config", "./config.yml", "The path to the configuration file")
	flag.Parse()

	err := config.Load(*cpath)
	if err != nil {
		log.Fatalf("CRITICAL %s", err)
	}
	scheduler.Start()
	server.Serve()
}
