package articlestore

import (
	"github.com/GaryBrownEEngr/twertle_api_dev/backend/models"
)

type store struct {
	Articles []models.Article
	NextId   int
}

func NewStore(articles []models.Article) *store {
	s := store{Articles: articles, NextId: articles[len(articles)-1].Id + 1}
	return &s
}

func (a *store) CreateArticle(newArticle *models.Article) {
	newArticle.Id = a.NextId
	a.NextId++
	a.Articles = append(a.Articles, *newArticle)
}

func (a *store) GetArticle(id int) *models.Article {
	for _, article := range a.Articles {
		if article.Id == id {
			return &article
		}
	}
	return nil
}

func (a *store) GetArticles() []models.Article {
	return a.Articles
}

func (a *store) DeleteArticle(id int) {
	var foundIndex int = -1
	for i, article := range a.Articles {
		if article.Id == id {
			foundIndex = i
		}
	}

	if foundIndex != -1 {
		a.Articles = append(a.Articles[:foundIndex], a.Articles[foundIndex+1:]...)
	}
}
