package helpscout

import (
	"fmt"
)

type SummaryArticle struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	IsPublic bool   `json:"public"`
}

type SummaryCategory struct {
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	URL         string            `json:"url"`
	IsPublic    bool              `json:"public"`
	Items       []*SummaryArticle `json:"items"`
}

type SummaryCollection struct {
	Id          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	URL         string             `json:"url"`
	IsPublic    bool               `json:"public"`
	Items       []*SummaryCategory `json:"items"`
}

type Summary struct {
	RateLimitRemain int                  `json:"ratelimit_remain"`
	RateLimitReset  int                  `json:"ratelimit_reset"`
	Items           []*SummaryCollection `json:"items"`
}

func (c *SummaryCategory) getAllArticles(hs *HelpScout) error {
	page := 1
	pages := 1

	for page <= pages {
		articles, err := hs.ListArticles(c.Id, page)
		if err != nil {
			return fmt.Errorf("receiving article from %s: %w", c.Name, err)
		}

		page += 1
		pages = articles.Pages

		for _, item := range articles.Items {
			article := &SummaryArticle{
				Id:       item.Id,
				Name:     item.Name,
				URL:      item.PublicUrl,
				IsPublic: item.Status == "published",
			}

			c.Items = append(c.Items, article)
		}
	}

	return nil
}

func (c *SummaryCollection) getAllCategories(hs *HelpScout) error {
	page := 1
	pages := 1

	for page <= pages {
		categories, err := hs.ListCategories(c.Id, page)
		if err != nil {
			return fmt.Errorf("receiving categories from %s: %w", c.Name, err)
		}

		page += 1
		pages = categories.Pages

		for _, item := range categories.Items {
			category := &SummaryCategory{
				Id:          item.Id,
				Name:        item.Name,
				Description: item.Description,
				URL:         item.PublicUrl,
				IsPublic:    item.Visibility == "public",
			}

			c.Items = append(c.Items, category)

			if item.ArticleCount == 0 {
				continue
			}

			if err := category.getAllArticles(hs); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Summary) getAllCollections(hs *HelpScout) error {
	page := 1
	pages := 1

	for page <= pages {
		collections, err := hs.ListCollections(page)
		if err != nil {
			return fmt.Errorf("receiving collections: %w", err)
		}

		page += 1
		pages = collections.Pages

		for _, item := range collections.Items {
			collection := &SummaryCollection{
				Id:          item.Id,
				Name:        item.Name,
				Description: item.Description,
				URL:         item.PublicUrl,
				IsPublic:    item.Visibility == "public",
			}

			s.Items = append(s.Items, collection)

			if err := collection.getAllCategories(hs); err != nil {
				return err
			}
		}
	}

	return nil
}

func (hs *HelpScout) GetSummary() (*Summary, error) {
	summary := &Summary{}

	if err := summary.getAllCollections(hs); err != nil {
		return nil, err
	}

	summary.RateLimitRemain = hs.rateLimitRemain
	summary.RateLimitReset = hs.rateLimitReset

	return summary, nil
}
