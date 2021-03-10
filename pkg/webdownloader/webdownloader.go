package webdownloader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/bend-is/web-page-downloader/pkg/webdownloader/wordprocessor"
)

type WebDownloader struct {
	client       *http.Client
	wordAnalyser *wordprocessor.WordProcessor
}

type WebPageStatistic struct {
	UniqueWordCount map[string]int
}

func New(client *http.Client) *WebDownloader {
	if client == nil {
		client = http.DefaultClient
	}

	return &WebDownloader{
		client:       client,
		wordAnalyser: wordprocessor.New(),
	}
}

func (pd *WebDownloader) Download(targetUrl string) (*WebPageStatistic, error) {
	u, err := url.ParseRequestURI(targetUrl)
	if err != nil {
		return nil, err
	}

	res, err := pd.client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	f, err := os.OpenFile(pd.getFileName(u), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if _, err = io.Copy(f, res.Body); err != nil {
		return nil, err
	}
	if _, err = f.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	return &WebPageStatistic{pd.wordAnalyser.UniqueWordsCount(f)}, nil
}

func (pd *WebDownloader) getFileName(u *url.URL) string {
	return strings.Replace(fmt.Sprintf("%s%s.html", u.Host, u.Path), "/", "", -1)
}
