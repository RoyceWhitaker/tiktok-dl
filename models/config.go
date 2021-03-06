package models

import (
	"flag"
	"fmt"
	"os"
)

// Config - Runtime configuration
var Config struct {
	URL        string
	OutputPath string
	Debug      bool
	MetaData   bool
}

// GetConfig - Returns Config object
func GetConfig() {
	outputPath := flag.String("output", "./downloads", "Output path")
	debug := flag.Bool("debug", false, "Enables debug mode")
	metadata := flag.Bool("metadata", false, "Write video metadata to a .json file")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: tiktok-dl [OPTIONS] TIKTOK_USERNAME|TIKTOK_URL")
		os.Exit(2)
	}

	Config.URL = flag.Args()[len(args)-1]
	Config.OutputPath = *outputPath
	Config.Debug = *debug
	Config.MetaData = *metadata
}
