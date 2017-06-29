package main

import (
	"fmt"

	"github.com/gnur/golpje/episodes"
)

func main() {
	title := "Ridiculousness.S09E27.Gene.Dyrdek.576p..WEBRip.AAC2.0.H.264-BTW"

	a, b := episodes.ExtractEpisodeID(title, true)
	fmt.Println(a)
	fmt.Println(b)

}
