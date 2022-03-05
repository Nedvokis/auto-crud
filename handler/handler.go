package handler

import "github.com/auto-crud/db"

type Handler struct {
	store db.Store
}

func NewHandler(store *db.Store) *Handler {
	return &Handler{
		store: *store,
	}
}
