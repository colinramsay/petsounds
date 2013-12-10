package main

import (
    "log"
	"net/http"
	"html/template"
    //"database/sql"
    //
    "github.com/colinramsay/go-musicbrainz"
)

func renderTemplate(w http.ResponseWriter, tmpl string, result interface{}) {
    t, _ := template.ParseFiles("./tpl/" + tmpl + ".html")
    t.Execute(w, result)
}

func rootHandler(w http.ResponseWriter, r *http.Request) { 
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

func main() {
	// db, err := sql.Open("sqlite3", "./petsounds.db")
 //    if err != nil {
 //        log.Fatal(err)
 //    }
 //    defer db.Close()


    mux := http.NewServeMux()
    mux.Handle("/", http.HandlerFunc(rootHandler))
    mux.Handle("/search", http.HandlerFunc(artistSearchHandler))
    mux.Handle("/releases", http.HandlerFunc(releasesHandler))

    mux.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
        log.Printf("serving %s", "./public"+r.URL.Path)
        http.ServeFile(w, r, "./public"+r.URL.Path)

    })

    http.ListenAndServe(":8999", mux)
}