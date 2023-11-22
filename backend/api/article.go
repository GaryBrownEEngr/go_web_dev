package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils"

	"github.com/gorilla/mux"
)

func returnAllArticles(articles models.ArticleStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint Hit: returnAllArticles")
		utils.EncodeJSON(w, articles.GetArticles(), http.StatusOK)
	}
}

func returnSingleArticle(articles models.ArticleStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.EncodeJSON(w, apiError{"Id param is invalid"}, http.StatusBadRequest)
			return
		}

		a := articles.GetArticle(key)
		if a == nil {
			utils.EncodeJSON(w, apiError{"article not found"}, http.StatusBadRequest)
			return
		}

		utils.EncodeJSON(w, a, http.StatusOK)
	}
}

func deleteSingleArticle(articles models.ArticleStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.EncodeJSON(w, apiError{"Id param is invalid"}, http.StatusBadRequest)
			return
		}

		articles.DeleteArticle(key)

		utils.EncodeJSON(w, nil, http.StatusOK)
	}
}

func createNewArticle(articles models.ArticleStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newArticle models.Article
		if err := json.NewDecoder(r.Body).Decode(&newArticle); err != nil {
			utils.EncodeJSON(w, apiError{"unable to parse JSON"}, http.StatusBadRequest)
			return
		}
		articles.CreateArticle(&newArticle)
		utils.EncodeJSON(w, newArticle, http.StatusCreated)
	}
}
