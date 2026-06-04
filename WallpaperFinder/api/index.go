package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type SrcSizes struct {
	Original string `json:"original"`
	Medium   string `json:"medium"`
	Large2x  string `json:"large2x"`
}

type Wallpaper struct {
	Id           int      `json:"id"`
	Photographer string   `json:"photographer"`
	Src          SrcSizes `json:"src"`
	Url          string   `json:"url"`
}

type PexelsResponse struct {
	Photos []Wallpaper `json:"photos"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("query")
	page := r.URL.Query().Get("page")
	if query == "" {
		query = "curated"
	}
	if page == "" {
		page = "1"
	}
	apiKey := os.Getenv("PEXELS_API_KEY")

	url := fmt.Sprintf("https://api.pexels.com/v1/search?query=%s&per_page=12&page=%s", query, page)
	if query == "curated" {
		url = fmt.Sprintf("https://api.pexels.com/v1/curated?per_page=12&page=%s", page)
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"photos":[]}`))
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
