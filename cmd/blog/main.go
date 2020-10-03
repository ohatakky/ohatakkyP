package main

import (
	"fmt"

	"github.com/ohatakky/ohatakkyp/pkg/rss"
)

var (
	urls = []string{
		"https://future-architect.github.io/atom.xml",
		"https://buildersbox.corp-sansan.com/rss",
	}
)

func main() {
	reader := rss.New()
	feeds := reader.Read(urls)
	fmt.Println(feeds)
}
