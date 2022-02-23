package models

type Payload struct {
	URL  string `json:"url"`
	Slug string `json:"slug"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ShortURL struct {
	Short string `json:"short_url"`
}
