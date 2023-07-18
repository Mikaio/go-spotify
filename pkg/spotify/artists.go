package spotify

type ExternalUrls struct {
	Spotify string `json:"spotify"`
}

type Followers struct {
	Href  string `json:"href"`
	Total uint32 `json:"total"`
}

type Genres []string

type Image struct {
	Url    string `json:"url"`
	Height uint16 `json:"height"`
	Width  uint16 `json:"width"`
}

type Artist struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Followers    Followers    `json:"followers"`
	Genres       Genres       `json:"genres"`
	Href         string       `json:"href"`
	Id           string       `json:"id"`
	Images       []Image      `json:"images"`
	Name         string       `json:"name"`
	Popularity   uint16       `json:"popularity"`
	Type         string       `json:"type"`
	Uri          string       `json:"uri"`
}
