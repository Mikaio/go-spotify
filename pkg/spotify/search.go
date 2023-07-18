package spotify

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
)

type Search struct {
	Href     string  `json:"hred"`
	Limit    uint    `json:"limit"`
	Next     string  `json:"next"`
	Offset   uint    `json:"offset"`
	Previous string  `json:"previous"`
	Total    uint    `json:"total"`
	Items    []Track `json:"items"`
}

type TracksSearch struct {
	Tracks Search `json:"tracks"`
}

type SearchType string

const (
	ALBUM     SearchType = "album"
	ARTIST    SearchType = "artist"
	TRACK     SearchType = "track"
	SHOW      SearchType = "show"
	PLAYLIST  SearchType = "playlist"
	EPISODE   SearchType = "episode"
	AUDIOBOOK SearchType = "audiobook"
)

type SearchProps struct {
	Query  string
	Type   SearchType
	Limit  uint
	Offset uint
}

func (s *Spotify) Search(searchProps *SearchProps) (*TracksSearch, error) {

	data := url.Values{}
	data.Set("q", searchProps.Query)
	data.Set("limit", fmt.Sprint(searchProps.Limit))
	data.Set("offset", fmt.Sprint(searchProps.Offset))
	data.Set("type", string(searchProps.Type))

	resource := "search"

	reqUrl, err := url.ParseRequestURI(s.ApiV1)

	if err != nil {
		log.Fatal("failed to parse api url")
		return nil, errors.New("could not search")
	}

	reqUrl.Path = path.Join(reqUrl.Path, resource)
	reqUrl.RawQuery = data.Encode()

	client := http.Client{}

	fmt.Println(reqUrl)

	req, err := http.NewRequest("GET", fmt.Sprint(reqUrl), nil)

	if err != nil {
		log.Fatal("failed to create request instance")
		return nil, errors.New("could not search")
	}

	authorizationHeader := "Bearer " + s.AccessToken
	req.Header.Add("Authorization", authorizationHeader)

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("could not search")
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal("could not search. Status: ", res.StatusCode)
		return nil, errors.New("could not search")
	}

	var tracksSearch TracksSearch
	err = json.NewDecoder(res.Body).Decode(&tracksSearch)

	if err != nil {
		log.Fatal("could not extract search result")
		return nil, err
	}

	return &tracksSearch, nil
}
