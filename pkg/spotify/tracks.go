package spotify

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"strings"

	"log"
	"net/http"
	"net/url"
)

type LinkedFrom struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	Id           string       `json:"id"`
	Type         string       `json:"type"`
	Uri          string       `json:"uri"`
}

type Track struct {
	Album            Album        `json:"album"`
	Artists          []Artist     `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       uint16       `json:"disc_number"`
	DurationMs       uint32       `json:"duration_ms"`
	ExternalIds      ExternalIds  `json:"external_ids"`
	ExternalUrls     ExternalUrls `json:"external_urls"`
	Href             string       `json:"href"`
	Id               string       `json:"id"`
	IsPlayable       bool         `json:"is_playable"`
	LinkedFrom       LinkedFrom   `json:"linked_from"`
	Restrictions     Restrictions `json:"restrictions"`
	Name             string       `json:"name"`
	Popularity       uint8        `json:"popularity"`
	PreviewUrl       string       `json:"preview_url"`
	TrackNumber      uint16       `json:"track_number"`
	Type             string       `json:"type"`
	Uri              string       `json:"uri"`
	IsLocal          bool         `json:"is_local"`
}

type Seed struct {
	InitialPoolSize    uint   `json:"initialPoolSize"`
	AfterFilteringSize uint   `json:"afterFilteringSize"`
	AfterRelinkingSize uint   `json:"afterRelinkingSize"`
	Id                 string `json:"id"`
	Type               string `json:"type"`
	Href               string `json:"href"`
}

type TrackRecommendations struct {
	Tracks []Track `json:"tracks"`
	Seeds  []Seed  `json:"seeds"`
}

type SpotifyError struct {
	Message string
	Err     error
}

func (s *Spotify) GetTrackInfo(trackId string) (*Track, error) {

	client := http.Client{}

	resource := "/tracks/" + trackId

	authorizationHeader := "Bearer " + s.AccessToken

	reqUrl, err := url.JoinPath(s.ApiV1, resource)

	if err != nil {
		log.Fatal("COULD NOT PARSE REQUEST URL:\n", err)
	}

	req, err := http.NewRequest("GET", reqUrl, nil)

	if err != nil {
		log.Fatal("ERROR CREATING TRACK INFO REQUEST INSTANCE:\n", err)
	}

	req.Header.Add("Authorization", authorizationHeader)

	res, err := client.Do(req)

	if err != nil {
		log.Fatal("ERROR GETTING TRACK:\n", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Error getting track with status " + fmt.Sprint(res.StatusCode))
	}

	var result Track

	err = json.NewDecoder(res.Body).Decode(&result)

	if err != nil {
		panic("Could not decode track result")
	}

	println("TRACK NAME: " + result.Name)
	println("URL: " + result.ExternalUrls.Spotify)
	println("ALBUM: " + result.Album.ExternalUrls.Spotify)

	fmt.Printf("%+v\n", result)

	return &result, nil
}

type GetTrackRecommendationsProps struct {
	Limit       uint8
	SeedTracks  []string
	SeedArtists []string
	SeedGenres  []string
}

func (tr GetTrackRecommendationsProps) ExtractTracks() (string, error) {
	stIds := make([]string, len(tr.SeedTracks))

	for track := range tr.SeedTracks {
		trackSplit := strings.Split(tr.SeedTracks[track], "/track/")

		fmt.Println("track split: ", trackSplit)

		if len(trackSplit) < 2 {
			return "", errors.New("could not extract track id")
		}

		querySplit := strings.Split(trackSplit[1], "?")

		fmt.Println("query split: ", querySplit)

		if len(querySplit) < 1 {
			return "", errors.New("could not extract track id")
		}

		stIds[track] = querySplit[0]
	}

	fmt.Println(stIds)

	return strings.Join(stIds, ","), nil
}

func (s *Spotify) GetRecommendations(trackRecommendationProps *GetTrackRecommendationsProps) (*TrackRecommendations, error) {
	trackIds, err := trackRecommendationProps.ExtractTracks()
	fmt.Println("limit:", fmt.Sprint(trackRecommendationProps.Limit))

	if err != nil {
		panic(err)
	}

	fmt.Println(trackIds)

	data := url.Values{}
	data.Set("limit", fmt.Sprint(trackRecommendationProps.Limit))
	data.Set("seed_artists", "")
	data.Set("seed_genres", "")
	data.Set("seed_tracks", trackIds)

	resource := "recommendations"

	reqUrl, err := url.ParseRequestURI(s.ApiV1)

	if err != nil {
		panic("bruh")
	}

	reqUrl.Path = path.Join(reqUrl.Path, resource)
	reqUrl.RawQuery = data.Encode()

	fmt.Println(reqUrl)

	client := http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprint(reqUrl), nil)

	if err != nil {
		log.Fatal("error creating request instance")
		return nil, errors.New("could not get recommendations")
	}

	authorizationHeader := "Bearer " + s.AccessToken
	req.Header.Add("Authorization", authorizationHeader)

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("could not get recommendations")
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal("could not get tracks recommendations. Status: ", res.StatusCode)
		return nil, errors.New("could not get recommendations")
	}

	var result TrackRecommendations

	err = json.NewDecoder(res.Body).Decode(&result)

	if err != nil {
		log.Fatal("could not extract track recommendations")
		return nil, errors.New("could not get recommendations")
	}

	return &result, nil
}
