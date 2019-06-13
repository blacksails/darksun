package main

import (
	"github.com/blacksails/darksun/iterm2"
	"github.com/blacksails/darksun/macos"
)

type config struct {
	Modules modules
}

type modules struct {
	ITerm2 iterm2.Config
	MacOS  macos.Config
}
