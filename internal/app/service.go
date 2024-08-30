package app

import (
	"Calorie_calculator/internal/config"
	"Calorie_calculator/internal/pkg/store"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type Service struct {
	Port   string
	router *mux.Router
	store  store.Store
}

func NewService(db *sql.DB) Service {
	tmp := Service{
		Port:   config.ServPort,
		router: mux.NewRouter(),
		store:  store.NewStore(db),
	}
	tmp.SetHandlers()
	return tmp
}

func (s Service) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func (s Service) SetHandlers() {
	api := s.router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/calc", s.CalculatorFunc) //Добавление новых записей и запрос всей таблицы

	api.HandleFunc("/calc_prod", s.CalcProdFunc)
}
