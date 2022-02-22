package models

type Payload struct {
	URL  string `json:"url"`
	Slug string `json:"slug"`
}

type Error struct {
	Message string `json:"error"`
}

type ShortURL struct {
	Short string `json:"short_url"`
}
