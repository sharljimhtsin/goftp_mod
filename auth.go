// Copyright 2018 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Auth is an interface to auth your ftp user login.
type Auth interface {
	CheckPasswd(string, string) (bool, error)
}

var (
	_ Auth = &SimpleAuth{}
)

// SimpleAuth implements Auth interface to provide a memory user login auth
type SimpleAuth struct {
	Cfg string
}

// CheckPasswd will check user's password
func (a *SimpleAuth) CheckPasswd(name, pass string) (bool, error) {
	cfgPath := a.Cfg
	file, err := os.Open(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	success := false
	for scanner.Scan() {
		txt := scanner.Text()
		args := strings.Split(txt, ":")
		if name == args[0] && pass == args[1] {
			success = true
			break
		}
	}
	return success, nil
}
