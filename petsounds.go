package main

import (
	"encoding/json"
	"fmt"
	"github.com/colinramsay/go-musicbrainz"
	scrapers "github.com/colinramsay/petsounds_scrapers"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"flag"
)

func saveSettingsHandler(w http.ResponseWriter, r *http.Request) {
	settings := Settings {
		TorrentConfiguration {
			r.FormValue("torrentBlackHole"),
			r.FormValue("pirateBayProxyUrl"),
			"",
		},
		"",
	}

	bytes, err := json.MarshalIndent(settings, "", "    ")

	err = ioutil.WriteFile(CONFIG_FILE, bytes, 0644)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Saved settings!")
}

func readSettings() Settings {
	bytes, err := ioutil.ReadFile(CONFIG_FILE)

	if err != nil {
		panic("Could not find configuration file.")
	}

	var settings Settings

	err = json.Unmarshal(bytes, &settings)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded configuration file %v", settings)

	return settings
}

func showSettingsHandler(w http.ResponseWriter, r *http.Request) {
	settings := readSettings()

	renderTemplate(w, "settings", settings)
}

func renderTemplate(w http.ResponseWriter, tmpl string, result interface{}) {
	t, _ := template.ParseFiles("./tpl/" + tmpl + ".html")
	t.Execute(w, result)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "custom 404")
		return
	}
	renderTemplate(w, "index", nil)
}

func releasesHandler(w http.ResponseWriter, r *http.Request) {
	result := musicbrainz.ReleaseResult{}
	result = musicbrainz.GetReleases(r.FormValue("id"))
	renderTemplate(w, "releases", result)
}

func artistSearchHandler(w http.ResponseWriter, r *http.Request) {
	result := musicbrainz.ArtistResult{}

	if r.Method == "POST" {
		result = musicbrainz.SearchArtist(r.FormValue("artist"))
	}

	renderTemplate(w, "artist", result)
}

func fetchHandler(w http.ResponseWriter, r *http.Request) {
	settings := readSettings()

	pb := scrapers.NewPirateBay(settings.TorrentConfiguration.PirateBayProxy)
	term := r.FormValue("term")
	filename := pb.SearchAndSave(term, settings.TorrentConfiguration.BlackHoleDirectory)

	fmt.Fprintf(w, "File fetched to %s", filename)
}

var CONFIG_FILE string

func main() {

	flag.StringVar(&CONFIG_FILE, "config", "./petsounds.conf.json", "Path to the config file")
	flag.Parse()

	log.Printf("Using config from %v", CONFIG_FILE)

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(rootHandler))
	mux.Handle("/search", http.HandlerFunc(artistSearchHandler))
	mux.Handle("/releases", http.HandlerFunc(releasesHandler))
	mux.Handle("/release/fetch", http.HandlerFunc(fetchHandler))

	mux.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			saveSettingsHandler(w, r)
		} else if r.Method == "GET" {
			showSettingsHandler(w, r)
		}
	})

	mux.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("serving %s", "./public"+r.URL.Path)
		http.ServeFile(w, r, "./public"+r.URL.Path)

	})

	go http.ListenAndServe(":8999", mux)
	settings := readSettings()
	pp := PostProcessor{}
	pp.StartWatching(settings.TorrentConfiguration.CompleteDirectory, settings.PostProcessingScript)

}
