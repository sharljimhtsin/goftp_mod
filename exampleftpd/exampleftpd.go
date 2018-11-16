// Copyright 2018 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// This is a very simple ftpd server using this library as an example
// and as something to run tests against.
package main

import (
	"flag"
	"log"
	"server"
	"server/file-driver"
)

func main() {
	var (
		root = flag.String("root", "", "Root directory to serve")
		port = flag.Int("port", 2121, "Port")
		host = flag.String("host", "localhost", "Port")
		cfg  = flag.String("cfg", "", "path to cfg")
	)
	flag.Parse()
	if *root == "" {
		log.Fatalf("Please set a root to serve with -root")
	}

	if *cfg == "" {
		log.Fatalf("Please set a cfg to serve with -cfg")
	}

	factory := &filedriver.FileDriverFactory{
		RootPath: *root,
		Perm:     server.NewSimplePerm("user", "group"),
	}

	opts := &server.ServerOpts{
		Factory:  factory,
		Port:     *port,
		Hostname: *host,
		Auth:     &server.SimpleAuth{Cfg: *cfg},
	}

	log.Printf("Starting ftp server on %v:%v", opts.Hostname, opts.Port)
	log.Printf("cfg path is %v", *cfg)
	server := server.NewServer(opts)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
