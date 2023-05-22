package response

type PostResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `jsosn:"description"`
}
