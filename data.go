package main

type TorrentConfiguration struct {
	PirateBayProxy string
	BlackHoleDirectory string
}

type Settings struct {
	TorrentConfiguration TorrentConfiguration
}