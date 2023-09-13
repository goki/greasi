// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"goki.dev/grease"
	"goki.dev/greasi"
)

//go:generate gtigen

type Config struct {

	// the name of the user
	Name string `desc:"the name of the user"`

	// the age of the user
	Age int `desc:"the age of the user"`

	// whether the user likes Go
	LikesGo bool `desc:"whether the user likes Go"`

	// the target platform to build for
	BuildTarget string `desc:"the target platform to build for"`
}

// Build builds the app for the config build target.
//
//gti:add
func Build(c *Config) error {
	fmt.Println("Building for platform", c.BuildTarget)
	return nil
}

// Run runs the app for the user with the config name.
//
//gti:add
func Run(c *Config) error {
	fmt.Println("Running for user", c.Name)
	return nil
}

func main() {
	grease.AppName = "basic"
	grease.AppTitle = "Basic"
	grease.AppAbout = "Basic is a basic example application made with Grease."
	grease.DefaultFiles = []string{"config.toml"}
	greasi.Run(&Config{}, Build, Run)
}