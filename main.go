package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	country  = flag.String("country", "", "set country query parameter")
	category = flag.String("category", "", "set category query parameter")
	keywords = flag.String("q", "", "set keywords query parameter")
	sources  = flag.String("sources", "", "set sources query parameter")
	language = flag.String("language", "en", "set language query parameter")
)

func main() {

	cfg := GetConfig()

	flag.Parse()

	query := Query{
		Category: *category,
		Country:  *country,
		Keywords: *keywords,
		Sources:  *sources,
		Language: *language,
	}

	getNews(cfg, query)
}

func createAPIQuery(cfg *Config, query Query) string {

	var APIurl = "https://newsapi.org/v2/top-headlines"
	if cfg.APIurl != "" {
		APIurl = cfg.APIurl
	}

	values, err := query.Values()
	if err != nil {
		log.Fatal(err)
	}
	values.Add("apiKey", cfg.APIkey)

	endPoint, err := url.Parse(APIurl)
	if err != nil {
		log.Fatal(err)
	}

	endPoint.RawQuery = values.Encode()

	return endPoint.String()
}

func getNews(cfg *Config, query Query) {

	resp, err := http.Get(createAPIQuery(cfg, query))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	response := &Response{}
	if err := json.Unmarshal(body, response); err != nil {
		log.Fatal(err)
	}

	showNews(response)

}

func showNews(response *Response) {
	if response.Status == "error" {
		return
	}

	tmpl, err := template.New("article").Parse(
		`[{{.PublishedAt.Format "2006-01-02 15:04"}}] {{.Title}} by {{.Author}}
{{.Description}}
URL: {{.URL}}

`)
	if err != nil {
		log.Fatal(err)
	}

	for _, article := range response.Articles {
		err = tmpl.Execute(os.Stdout, article)
		if err != nil {
			log.Fatal(err)
		}
	}
}
