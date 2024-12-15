package main

import (
	"os"

	"github.com/gopy-art/zrediss/bin"
	"github.com/gopy-art/zrediss/console"
	logger "github.com/gopy-art/zrediss/log"
)

func init() {
	console.InitFlags()

	// init logger
	if console.Logger == "stdout" {
		logger.InitLoggerStdout()
	} else if console.Logger == "file" {
		zLabBaseDir, _ := os.Getwd()
		logger.InitLoggerFile(zLabBaseDir + "/zlab.log")
	}
}

func main() {
	bin.RunModuleWithFlag()
}
