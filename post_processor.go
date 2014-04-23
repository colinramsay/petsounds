package main

import (
	"github.com/howeyc/fsnotify"
	"log"
	"os/exec"
)

type PostProcessor struct {
}

func (pp *PostProcessor) StartWatching(directory string, shellScript string) {
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
				log.Println("Watch event: ", ev)

				err := exec.Command(shellScript, ev.Name).Run()
				if err != nil {
					log.Fatal(err)
				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.WatchFlags(directory, fsnotify.FSN_CREATE)
	if err != nil {
		log.Fatal(err)
	}

	<-done

	watcher.Close()
}
