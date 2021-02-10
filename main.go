package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

func main() {
	aclient := search.NewClient(os.Getenv("ALGOLIA_APPID"), os.Getenv("ALGOLIA_SEARCH_API_KEY"))
	index := aclient.InitIndex("shakespeare_v1")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/search", handleSearch(&Algolia{index}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("Listening on port %s...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Searcher interface {
	Search(string) ([]Result, error)
}

type Result struct {
	Play    string
	Speaker string
	Text    string
}

func handleSearch(searcher Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}
		results, err := searcher.Search(query[0])
		if err != nil {
			log.Println("q="+query[0], "err="+err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(results)
		if err != nil {
			log.Println("q="+query[0], "err="+err.Error())
		}
	}
}

type Algolia struct {
	index *search.Index
}

func (a *Algolia) Search(q string) ([]Result, error) {
	r, err := a.index.Search(q, opt.HitsPerPage(1000))
	if err != nil {
		return nil, fmt.Errorf("algolia search: %w", err)
	}
	var hits []Hit
	err = r.UnmarshalHits(&hits)
	if err != nil {
		return nil, fmt.Errorf("unmarshal hits: %w", err)
	}
	results := make([]Result, 0, len(r.Hits))
	for _, h := range hits {
		results = append(results, Result{
			Play:    h.HighlightResult.PlayName.Value,
			Speaker: h.HighlightResult.Speaker.Value,
			Text:    h.HighlightResult.TextEntry.Value,
		})
	}
	return results, nil
}

type Hit struct {
	Type          string `json:"type"`
	LineID        int    `json:"line_id"`
	PlayName      string `json:"play_name"`
	SpeechNumber  int    `json:"speech_number"`
	LineNumber    string `json:"line_number"`
	Speaker       string `json:"speaker"`
	TextEntry     string `json:"text_entry"`
	ObjectID      string `json:"objectID"`
	SnippetResult struct {
		Type struct {
			Value      string `json:"value"`
			MatchLevel string `json:"matchLevel"`
		} `json:"type"`
		LineID struct {
			Value      string `json:"value"`
			MatchLevel string `json:"matchLevel"`
		} `json:"line_id"`
		PlayName struct {
			Value      string `json:"value"`
			MatchLevel string `json:"matchLevel"`
		} `json:"play_name"`
		SpeechNumber struct {
			Value      string `json:"value"`
			MatchLevel string `json:"matchLevel"`
		} `json:"speech_number"`
		LineNumber struct {
			Value      string `json:"value"`
			MatchLevel string `json:"matchLevel"`
		} `json:"line_number"`
		Speaker struct {
			Value      string `json:"value"`
			MatchLevel string `json:"matchLevel"`
		} `json:"speaker"`
		TextEntry struct {
			Value      string `json:"value"`
			MatchLevel string `json:"matchLevel"`
		} `json:"text_entry"`
	} `json:"_snippetResult"`
	HighlightResult struct {
		Type struct {
			Value        string        `json:"value"`
			MatchLevel   string        `json:"matchLevel"`
			MatchedWords []interface{} `json:"matchedWords"`
		} `json:"type"`
		LineID struct {
			Value        string        `json:"value"`
			MatchLevel   string        `json:"matchLevel"`
			MatchedWords []interface{} `json:"matchedWords"`
		} `json:"line_id"`
		PlayName struct {
			Value        string        `json:"value"`
			MatchLevel   string        `json:"matchLevel"`
			MatchedWords []interface{} `json:"matchedWords"`
		} `json:"play_name"`
		SpeechNumber struct {
			Value        string        `json:"value"`
			MatchLevel   string        `json:"matchLevel"`
			MatchedWords []interface{} `json:"matchedWords"`
		} `json:"speech_number"`
		LineNumber struct {
			Value        string        `json:"value"`
			MatchLevel   string        `json:"matchLevel"`
			MatchedWords []interface{} `json:"matchedWords"`
		} `json:"line_number"`
		Speaker struct {
			Value        string        `json:"value"`
			MatchLevel   string        `json:"matchLevel"`
			MatchedWords []interface{} `json:"matchedWords"`
		} `json:"speaker"`
		TextEntry struct {
			Value            string   `json:"value"`
			MatchLevel       string   `json:"matchLevel"`
			FullyHighlighted bool     `json:"fullyHighlighted"`
			MatchedWords     []string `json:"matchedWords"`
		} `json:"text_entry"`
	} `json:"_highlightResult"`
	RankingInfo struct {
		NbTypos           int `json:"nbTypos"`
		FirstMatchedWord  int `json:"firstMatchedWord"`
		ProximityDistance int `json:"proximityDistance"`
		UserScore         int `json:"userScore"`
		GeoDistance       int `json:"geoDistance"`
		GeoPrecision      int `json:"geoPrecision"`
		NbExactWords      int `json:"nbExactWords"`
		Words             int `json:"words"`
		Filters           int `json:"filters"`
	} `json:"_rankingInfo"`
}
