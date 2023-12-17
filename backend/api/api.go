package api

import (
	"net/http"
	"time"

	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils"

	"github.com/gorilla/mux"
)

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

type Server struct {
	articles   models.ArticleStore
	handler    http.Handler
	mux        *mux.Router
	secrets    models.SecretStore
	users      models.UserStore
	MathData   models.MathGameStore
	tokenMaker models.TokenMaker
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}

func NewServer(
	articles models.ArticleStore,
	secrets models.SecretStore,
	users models.UserStore,
	mathData models.MathGameStore,
	tokenMaker models.TokenMaker,
) *Server {
	myRouter := mux.NewRouter().StrictSlash(true)

	// Stack the middlewares logging, and CORS
	handler := utils.LogHttp(CORS(myRouter.ServeHTTP))

	ret := &Server{
		handler:    handler,
		mux:        myRouter,
		articles:   articles,
		secrets:    secrets,
		users:      users,
		MathData:   mathData,
		tokenMaker: tokenMaker,
	}

	ret.AddRoutes()

	return ret
}

func (s *Server) AddRoutes() {
	s.mux.HandleFunc("/api/articles", returnAllArticles(s.articles))
	s.mux.HandleFunc("/api/article", createNewArticle(s.articles)).Methods("POST")
	s.mux.HandleFunc("/api/article/{id}", returnSingleArticle(s.articles)).Methods("GET")
	s.mux.Handle("/api/article/{id}", deleteSingleArticle(s.articles)).Methods("DELETE")
	s.mux.HandleFunc("/api/checkguess", checkGuess)
	s.mux.Handle("/api/sleep", utils.TimeoutMiddleware(time.Second, sleep()))

	// User management
	s.mux.HandleFunc("/api/user/create", s.userCreate()).Methods("POST")
	s.mux.HandleFunc("/api/user/login", s.userLogin()).Methods("POST")
	s.mux.Handle("/api/user/delete", utils.EnsureAuth(s.tokenMaker, s.userDelete())).Methods("POST")

	// Math Game Data
	s.mux.Handle("/api/mathgame/read", utils.EnsureAuth(s.tokenMaker, s.mathGameRead())).Methods("GET")
	s.mux.Handle("/api/mathgame/write", utils.EnsureAuth(s.tokenMaker, s.mathGameWrite())).Methods("POST")

	fileServer := http.FileServer(http.Dir("./static"))
	s.mux.PathPrefix("/").Handler(fileServer)
}

type apiError struct {
	Error string `json:"error"`
}
