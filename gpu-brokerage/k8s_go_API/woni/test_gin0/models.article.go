// models.article.go

package main

import (
	"errors"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 원래는 DB  연결해야함 ! > 임시로 메로리에 저장하는걸
var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

//makeItem 삽입

//getAllItem 전체조회

//SearchItem 특정상품검색
