package models

type Article struct {
	Id      int    `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type ArticleStore interface {
	CreateArticle(newArticle *Article)
	GetArticle(id int) *Article
	GetArticles() []Article
	DeleteArticle(id int)
}
