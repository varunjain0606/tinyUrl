package handler

import (
	"github.com/varunjain0606/tinyUrl/src/api_service/server/cache"
	"github.com/varunjain0606/tinyUrl/src/api_service/server/database"
)

type UrlHandler struct {
	DB            database.Database
	Cache         cache.Store
}

func (u *UrlHandler) Authenticate(string) error {
	//panic("implement me")
	return nil
}

func (u *UrlHandler) CreateTinyUrl()  {
	//panic("implement me")

}

func (u *UrlHandler) GetOriginalUrl() {
	//return nil
}

func NewHandler(database database.Database, cache *cache.Store) (*UrlHandler, error) {
	return &UrlHandler{database, *cache }, nil
}

