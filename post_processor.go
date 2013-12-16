package main

import (
    "log"
    "github.com/howeyc/fsnotify"
)

type PostProcessor struct {
}

func (pp *PostProcessor) StartWatching(directory string) {
	watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }

    done := make(chan bool)

    // Process events
    go func() {
        for {
            select {
            case ev := <-watcher.Event:
                log.Println("event:", ev)
            case err := <-watcher.Error:
                log.Println("error:", err)
            }
        }
    }()

    err = watcher.Watch(directory)
    if err != nil {
        log.Fatal(err)
    }

    <-done

    watcher.Close()
}