package spotify

type AvailableMarkets []string

type Restrictions struct {
	Reason string `json:"reason"`
}

type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type ExternalIds struct {
	ISRC string `json:"isrc"`
	EAN  string `json:"ean"`
	UPC  string `json:"upc"`
}

type Album struct {
	AlbumType            string           `json:"album_type"`
	TotalTracks          uint16           `json:"total_tracks"`
	AvailableMarkets     AvailableMarkets `json:"available_markets"`
	ExternalUrls         ExternalUrls     `json:"external_urls"`
	Href                 string           `json:"href"`
	Id                   string           `json:"id"`
	Images               []Image          `json:"images"`
	Name                 string           `json:"name"`
	ReleaseDate          string           `json:"release_date"`
	ReleaseDatePrecision string           `json:"release_date_precision"`
	Restrictions         Restrictions     `json:"restrictions"`
	Type                 string           `json:"type"`
	Uri                  string           `json:"uri"`
	Copyrights           []Copyright      `json:"copyrights"`
	ExternalIds          ExternalIds      `json:"external_ids"`
	Genres               Genres           `json:"genres"`
	Label                string           `json:"label"`
	Popularity           uint8            `json:"popularity"`
	AlbumGroup           string           `json:"album_group"`
	Artists              []Artist         `json:"artists"`
}
