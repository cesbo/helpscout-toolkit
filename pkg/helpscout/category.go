package helpscout

import "time"

// Category is a group of articles
// Description: https://developer.helpscout.com/docs-api/objects/category/
// Endpoint: https://docsapi.helpscout.net/v1/categories/{id}
type Category struct {
	Id                    string    `json:"id"`
	Number                int       `json:"number"`
	Slug                  string    `json:"slug"`
	Visibility            string    `json:"visibility"`
	CollectionId          string    `json:"collectionId"`
	Order                 int       `json:"order"`
	DefaultSort           string    `json:"defaultSort"`
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
	ArticleCount          int       `json:"articleCount"`
	PublishedArticleCount int       `json:"publishedArticleCount"`
	PublicUrl             string    `json:"publicUrl"`
	CreatedBy             int       `json:"createdBy"`
	UpdatedBy             int       `json:"updatedBy"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}

// Categories is a list of categories
// Description: https://developer.helpscout.com/docs-api/categories/list/
// Endpoint: https://docsapi.helpscout.net/v1/collections/{id}/categories
type Categories struct {
	Page  int        `json:"page"`
	Pages int        `json:"pages"`
	Count int        `json:"count"`
	Items []Category `json:"items"`
}

type categoriesResponse struct {
	Categories *Categories `json:"categories"`
}

func ListCategories(collectionId string) (*Categories, error) {
	x := categoriesResponse{}
	path := "collections/" + collectionId + "/categories"
	err := getJSON(path, nil, &x)
	if err != nil {
		return nil, err
	}
	return x.Categories, nil
}
