package handler

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/mmcdole/gofeed"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path[1:]
	url, err := url.QueryUnescape(path)

	if url == "" {
		jsonErr := make(map[string]string)
		jsonErr["message"] = "No valid RSS URL found."
		jsonErr["documentation"] = "https://github.com/khalby786/rss-to-json"
		jsonResp, _ := json.Marshal(jsonErr)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)

		return
	}

	fp := gofeed.NewParser()
	feed, error := fp.ParseURL(url)

	if error != nil {
		jsonErr := make(map[string]string)
		jsonErr["message"] = error.Error()
		jsonErr["documentation"] = "https://github.com/khalby786/rss-to-json"
		jsonResp, _ := json.Marshal(jsonErr)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	jsonResp, err := json.Marshal(feed)
	if err != nil {
		jsonErr := make(map[string]string)
		jsonErr["message"] = err.Error()
		jsonErr["documentation"] = "https://github.com/khalby786/rss-to-json"
		jsonResp, _ := json.Marshal(jsonErr)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	} else {
		w.Write(jsonResp)
	}
}
