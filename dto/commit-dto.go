package dto

type Commit struct {
	Commit struct {
		Author struct {
			Name string `json:"name"`
			Date string `json:"date"`
		} `json:"author"`
		Message string `json:"message"`
	} `json:"commit"`
	Url string `json:"url"`
	Sha string `json:"sha"`
}
