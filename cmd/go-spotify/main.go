package main

import (
	"flag"
	"fmt"
	"log"

	spotifyService "github.com/mikaio/go-spotify/pkg/spotify"
)

var spotify *spotifyService.Spotify

var search string

func init() {
	spotify = &spotifyService.Spotify{
		ClientID:     SpotifyConfigs.ClientID,
		ClientSecret: SpotifyConfigs.ClientSecret,
		GrantType:    SpotifyConfigs.GrantType,
		TokenUrl:     SpotifyConfigs.TokenUrl,
		ApiV1:        SpotifyConfigs.ApiV1,
	}

	flag.StringVar(&search, "search", "", "search things")
	flag.Parse()
}

func main() {
	spotify.Authenticate()

	println("NEW ACCESS TOKEN: " + spotify.AccessToken)

	// track, err := spotify.GetTrackInfo("4pNx8RkUkN09rHZnE9II20")
	//
	// if err != nil {
	//     fmt.Println("Bruh")
	// }
	//
	// fmt.Println(track.Album.Name)

	// tracks := make([]string, 2)
	//
	// tracks[0] = "https://open.spotify.com/track/44UgawdeYuvQ1rKc33t2tp?si=739ced5b42a54553"
	// tracks[1] = "https://open.spotify.com/track/7grcN18xiQDGUVqWR7mUF0?si=92c7e24299154e31"
	//
	// trackRecommendationsProps := spotifyService.GetTrackRecommendationsProps{
	//     SeedTracks: tracks,
	//     Limit: 2,
	// }
	//
	// fmt.Println("getting recommendations")
	// trackRecommendations, err := spotify.GetRecommendations(&trackRecommendationsProps)
	//
	// if err != nil {
	//     log.Fatal(err)
	// }
	//
	// for track := range trackRecommendations.Tracks {
	//     fmt.Println("Name:", trackRecommendations.Tracks[track].Name)
	//     fmt.Println("Url:", trackRecommendations.Tracks[track].ExternalUrls.Spotify)
	//
	//     fmt.Println("")
	// }

	searchProps := spotifyService.SearchProps{
		Query:  search,
		Limit:  10,
		Offset: 1,
		Type:   spotifyService.TRACK,
	}

	searchResult, err := spotify.Search(&searchProps)

	if err != nil {
		log.Fatal("could not extract search result")
	}

	for index := range searchResult.Tracks.Items {
		fmt.Println("\nName: ", searchResult.Tracks.Items[index].Name)
		fmt.Println("Artist: ", searchResult.Tracks.Items[index].Artists[0].Name)
		fmt.Printf("Link: %v\n", searchResult.Tracks.Items[index].ExternalUrls.Spotify)
	}

}
