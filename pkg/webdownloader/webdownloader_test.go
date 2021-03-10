package webdownloader

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestWebDownloader_Download(t *testing.T) {
	expectedStat := map[string]int{
		"SEPTEMBER": 1,
		"IS":        1,
		"TIME":      1,
		"OF":        3,
		"BEGINNING": 3,
		"FOR":       1,
		"ALL":       1,
		"SCHOOL":    1,
		"FALL":      1,
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`September is a time
			Of beginning for all,
			Beginning of school
			Beginning of fall.`,
		))
	}))
	defer func() { testServer.Close() }()

	wd := New(http.DefaultClient)

	u, err := url.Parse(testServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(wd.getFileName(u))

	stat, err := wd.Download(testServer.URL)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range expectedStat {
		if count, exist := stat.UniqueWordCount[k]; !exist {
			t.Fatalf("Missing result key %s", k)
		} else if count != v {
			t.Fatalf("Wrong result count for key %s. Want %d got %d", k, v, count)
		}
	}
}
