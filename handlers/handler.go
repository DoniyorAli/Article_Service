package handlers

import (
	"UacademyGo/Blogpost/article_service/config"
	"UacademyGo/Blogpost/article_service/storage"
)

type handler struct {
	Stg storage.StorageInter
	Cfg config.Config
}

func NewHandler(stg storage.StorageInter, cfg config.Config) handler {
	return handler{
		Stg: stg,
		Cfg: cfg,
	}
}
