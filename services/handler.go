package services

import (
	"encoding/json"
	"homework_8/errs"
	"homework_8/storage"
	"io"
)

type handler struct {
	storage storage.Storage
}

func NewHandler(st *storage.Storage) *handler {
	return &handler{storage: *st}
}

func (h *handler) Add(item string, w io.Writer) error {
	err := h.storage.Add(item)

	if err != nil {
		if alreadyExistsErr, ok := err.(*errs.ErrAlreadyExists); ok {
			w.Write([]byte(alreadyExistsErr.GetDescription()))
			return nil
		}
		return err
	}
	return nil
}

func (h *handler) GetList(w io.Writer) error {

	usersList, err := h.storage.GetAll()

	if err != nil {
		return err
	}

	if len(usersList) > 0 {

		dat, err := json.Marshal(usersList)

		if err != nil {
			return err
		}
		w.Write(dat)
	}

	return nil
}

func (h *handler) GetUserById(id string, w io.Writer) error {

	user, err := h.storage.GetOne(id)

	if err != nil && !errs.ErrHasTypeNotFound(err) {
		return err
	}

	if err != nil {
		w.Write([]byte(""))
		return nil
	}

	dat, err := json.Marshal(user)

	if err != nil {
		return err
	}
	w.Write(dat)

	return nil
}

func (h *handler) Remove(id string, w io.Writer) error {

	err := h.storage.Delete(id)

	if err != nil {
		if notFoundErr, ok := err.(*errs.ErrNotFound); ok {
			w.Write([]byte(notFoundErr.GetDescription()))
			return nil
		}
		return err
	}

	return nil
}
