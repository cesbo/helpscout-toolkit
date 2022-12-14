package helpscout

import (
	"net/url"
	"strconv"
	"time"
)

// Collection is a group of categories
// Description: https://developer.helpscout.com/docs-api/objects/collection/
// Endpoint: https://docsapi.helpscout.net/v1/collections/{id}
type Collection struct {
	Id                    string    `json:"id"`
	SiteId                string    `json:"siteId"`
	Number                int       `json:"number"`
	Slug                  string    `json:"slug"`
	Visibility            string    `json:"visibility"`
	Order                 int       `json:"order"`
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
	PublicUrl             string    `json:"publicUrl"`
	ArticleCount          int       `json:"articleCount"`
	PublishedArticleCount int       `json:"publishedArticleCount"`
	CreatedBy             int       `json:"createdBy"`
	UpdatedBy             int       `json:"updatedBy"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}

// Collections is a list of collections
// Description: https://developer.helpscout.com/docs-api/collections/list/
// Endpoint: https://docsapi.helpscout.net/v1/collections
type Collections struct {
	Page  int          `json:"page"`
	Pages int          `json:"pages"`
	Count int          `json:"count"`
	Items []Collection `json:"items"`
}

type collectionsResponse struct {
	Collections *Collections `json:"collections"`
}

func (hs *HelpScout) ListCollections(page int) (*Collections, error) {
	x := collectionsResponse{}
	q := url.Values{
		"page": []string{strconv.Itoa(page)},
	}

	err := hs.getJSON("collections", q, &x)
	if err != nil {
		return nil, err
	}

	return x.Collections, nil
}
