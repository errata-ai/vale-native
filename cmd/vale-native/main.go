package main

import (
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/errata-ai/vale-native/internal/msg"
)

func init() {
	lf, err := newLogFile()
	if err != nil {
		panic(err)
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(lf, nil)))
}

func main() {
	cfg, err := readConfig()
	if err != nil {
		log.Fatal(err)
	} else if cfg.Path == "" {
		log.Fatal("no path to Vale binary found")
	}

	mgr, err := msg.NewStdioManager(cfg.Path)
	if err != nil {
		log.Fatal(err)
	}

	err = mgr.Run()
	if err == io.EOF {
		log.Println("received EOF; exiting...")
	} else if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
