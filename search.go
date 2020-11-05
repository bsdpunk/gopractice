package main

import (
	"fmt"

	"github.com/michiwend/gomusicbrainz"
)

func main() {

	// create a new WS2Client.
	client, _ := gomusicbrainz.NewWS2Client(
		"https://musicbrainz.org/ws/2",
		"A GoMusicBrainz example",
		"0.0.1-beta",
		"http://bsdpunk.blogspot.com",
	)

	// Search for some artist(s)
	resp, _ := client.SearchArtist(`artist:"Qveen Herby"`, -1, -1)
	var id gomusicbrainz.MBID
	// Pretty print Name and score of each returned artist.
	for _, artist := range resp.Artists {
		fmt.Printf("%+v\n", resp.WS2ListResponse)
		//fmt.Printf("Name: %-25sScore: %d\n", artist.Name, resp.Scores[artist])
		fmt.Printf("Name: %-25sScore: %d Id: %-25s\n", artist.Name, resp.Scores[artist], artist.ID)
		id = artist.ID
	}
	art, err := client.LookupArtist(id)
	//rel, err := client.LookupRelease(id)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", art)
	//fmt.Printf("%+v", rel)
}
