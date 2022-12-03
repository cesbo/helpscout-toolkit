package helpscout

import (
	"net/url"
	"strconv"
	"time"
)

// Description: https://developer.helpscout.com/docs-api/objects/article/
// Endpoint: https://docsapi.helpscout.net/v1/articles/{id}
type Article struct {
	Id           string    `json:"id"`
	Number       int       `json:"number"`
	CollectionId string    `json:"collectionId"`
	Slug         string    `json:"slug"`
	Status       string    `json:"status"`
	HasDraft     bool      `json:"hasDraft"`
	Name         string    `json:"name"`
	Text         string    `json:"text"`
	Categories   []string  `json:"categories"`
	Related      []string  `json:"related"`
	PublicUrl    string    `json:"publicUrl"`
	Popularity   float64   `json:"popularity"`
	ViewCount    int       `json:"viewCount"`
	CreatedBy    int       `json:"createdBy"`
	UpdatedBy    int       `json:"updatedBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	PublishedAt  time.Time `json:"lastPublishedAt"`
}

// Articles is a list of articles
// Description: https://developer.helpscout.com/docs-api/articles/list/
// Endpoint: https://docsapi.helpscout.net/v1/categories/{id}/articles
type Articles struct {
	Page  int        `json:"page"`
	Pages int        `json:"pages"`
	Count int        `json:"count"`
	Items []*Article `json:"items"`
}

type articlesResponse struct {
	Articles *Articles `json:"articles"`
}

func (hs *HelpScout) ListArticles(categoryId string, page int) (*Articles, error) {
	x := articlesResponse{}
	q := url.Values{
		"page": []string{strconv.Itoa(page)},
	}

	path := "categories/" + categoryId + "/articles"
	err := hs.getJSON(path, q, &x)
	if err != nil {
		return nil, err
	}

	return x.Articles, nil
}
