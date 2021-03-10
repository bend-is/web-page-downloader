package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bend-is/web-page-downloader/pkg/webdownloader"
)

func main() {
	enableFileLogging := flag.Bool("l", false, "enable logging to file")
	targetUrl := flag.String("u", "", "url for getting word statistic from")
	flag.Parse()

	if *targetUrl == "" {
		flag.PrintDefaults()
		os.Exit(2)
	}
	if *enableFileLogging {
		f, err := os.OpenFile(getLogFileName(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		log.SetOutput(io.Writer(f))
	}

	webDownloader := webdownloader.New(http.DefaultClient)

	stat, err := webDownloader.Download(*targetUrl)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	b, err := json.MarshalIndent(stat.UniqueWordCount, "", "  ")
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	b = append(b, '\n')

	if _, err = os.Stdout.Write(b); err != nil {
		log.Fatalf("ERROR: %s", err)
	}
}

func getLogFileName() string {
	return fmt.Sprintf("app_%s.log", time.Now().Format("02_01_2006"))
}
