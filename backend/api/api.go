package api

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/GaryBrownEEngr/go_web_dev/backend/models"

	"github.com/gorilla/mux"
)

type LogResponseWriter struct {
	http.ResponseWriter
	statusCode int
	buf        bytes.Buffer
}

func NewLogResponseWriter(w http.ResponseWriter) *LogResponseWriter {
	return &LogResponseWriter{ResponseWriter: w}
}

func (w *LogResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *LogResponseWriter) Write(body []byte) (int, error) {
	w.buf.Write(body)
	return w.ResponseWriter.Write(body)
}

type Server struct {
	articles   models.ArticleStore
	mux        *mux.Router
	secrets    models.SecretStore
	users      models.UserStore
	tokenMaker models.TokenMaker
}

func (m *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	logRespWriter := NewLogResponseWriter(w)

	//	... operation that takes 20 milliseconds ...
	//	t := time.Now()
	//	elapsed := t.Sub(start)
	logRespWriter.Header().Set("Access-Control-Allow-Origin", "*")
	logRespWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	logRespWriter.Header().Set("Access-Control-Allow-Methods", "*")

	if r.Method == "OPTIONS" {
		return
	}
	m.mux.ServeHTTP(logRespWriter, r)

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf(
		"duration=%dns status=%d path=%s\n",
		elapsed,
		logRespWriter.statusCode,
		r.URL.Path)
}

func NewServer(
	articles models.ArticleStore,
	secrets models.SecretStore,
	users models.UserStore,
	tokenMaker models.TokenMaker,
) *Server {
	myRouter := mux.NewRouter().StrictSlash(true)

	ret := &Server{
		mux:        myRouter,
		articles:   articles,
		secrets:    secrets,
		users:      users,
		tokenMaker: tokenMaker,
	}

	ret.AddRoutes()

	return ret
}

func (s *Server) AddRoutes() {
	s.mux.HandleFunc("/api/articles", returnAllArticles(s.articles))
	s.mux.HandleFunc("/api/article", createNewArticle(s.articles)).Methods("POST")
	s.mux.HandleFunc("/api/article/{id}", returnSingleArticle(s.articles)).Methods("GET")
	s.mux.HandleFunc("/api/article/{id}", deleteSingleArticle(s.articles)).Methods("DELETE")
	s.mux.HandleFunc("/api/checkguess", checkGuess)

	s.mux.HandleFunc("/api/user/create", s.userCreate()).Methods("POST")
	s.mux.HandleFunc("/api/user/login", s.userLogin()).Methods("POST")
	s.mux.HandleFunc("/api/user/delete", s.userDelete()).Methods("POST")

	fileServer := http.FileServer(http.Dir("./static"))
	s.mux.PathPrefix("/").Handler(fileServer)
}

type apiError struct {
	Error string `json:"error"`
}
