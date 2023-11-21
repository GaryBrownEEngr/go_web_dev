package api

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/GaryBrownEEngr/twertle_api_dev/backend/models"

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
	articles models.ArticleStore
	mux      *mux.Router
	secrets  models.SecretStore
	keyDB    models.KeyDBStore
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
	log.Printf(
		"duration=%dns status=%d path=%s",
		elapsed,
		logRespWriter.statusCode,
		r.URL.Path)
}

func NewServer(
	articles models.ArticleStore,
	secrets models.SecretStore,
	keyDB models.KeyDBStore,
) *Server {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/api/articles", returnAllArticles(articles))
	myRouter.HandleFunc("/api/article", createNewArticle(articles)).Methods("POST")
	myRouter.HandleFunc("/api/article/{id}", returnSingleArticle(articles)).Methods("GET")
	myRouter.HandleFunc("/api/article/{id}", deleteSingleArticle(articles)).Methods("DELETE")
	myRouter.HandleFunc("/api/checkguess", checkGuess)

	fileServer := http.FileServer(http.Dir("./static"))
	myRouter.PathPrefix("/").Handler(fileServer)

	return &Server{
		mux:      myRouter,
		articles: articles,
		secrets:  secrets,
		keyDB:    keyDB,
	}
}

type apiError struct {
	Error string `json:"error"`
}
