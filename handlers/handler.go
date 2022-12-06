package handlers

import (
	"UacademyGo/Article/config"
	"UacademyGo/Article/storage"
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