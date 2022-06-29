package storage

import (
	"encoding/json"
	"homework_8/errs"
	"homework_8/models"
	"io"
	"io/ioutil"
	"os"
)

var (
	user  = models.User{}
	users []models.User
)

type UserStorage interface {
	Add(item string) error
	GetOne(id string) (models.User, error)
	GetAll() (*[]models.User, error)
	Delete(id string) error
}

type Storage struct {
	Dir string
}

func NewStorage(dir string) *Storage {

	storage := &Storage{Dir: dir}

	return storage
}

func openConnection(dir string) (*os.File, error) {

	file, err := os.OpenFile(dir, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		os.Exit(1)
		return nil, err
	}

	return file, nil
}

func (s *Storage) GetAll() ([]models.User, error) {

	file, err := openConnection(s.Dir)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	dataFile, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	if len(dataFile) != 0 {

		err = json.Unmarshal(dataFile, &users)

		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

func (s *Storage) GetOne(id string) (models.User, error) {

	file, err := openConnection(s.Dir)

	if err != nil {
		return user, err
	}
	defer file.Close()

	dataFile, err := ioutil.ReadAll(file)

	if err != nil {
		return user, err
	}

	if len(dataFile) != 0 {

		err = json.Unmarshal(dataFile, &users)

		if err != nil {
			return user, err
		}
	}

	for _, user := range users {

		if user.ID == id {
			return user, nil
		}
	}

	return user, &errs.ErrNotFound{Id: id}
}

func (s *Storage) Delete(id string) error {

	file, err := openConnection(s.Dir)

	if err != nil {
		return err
	}
	defer file.Close()

	dataFile, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	if len(dataFile) != 0 {

		err = json.Unmarshal(dataFile, &users)

		if err != nil {
			return err
		}
	}

	for i, user := range users {

		if user.ID == id {

			users = append(users[:i], users[i+1:]...)

			dat, err := json.Marshal(users)

			if err != nil {
				return err
			}

			err = os.Truncate(s.Dir, 0)

			if err != nil {
				return err
			}

			file.Seek(0, io.SeekStart)
			file.Write(dat)

			return nil
		}
	}

	return &errs.ErrNotFound{Id: id}
}

func (s *Storage) Add(item string) error {

	file, err := openConnection(s.Dir)

	if err != nil {
		return err
	}
	defer file.Close()

	data := []byte(item)

	err = json.Unmarshal(data, &user)

	if err != nil {
		return err
	}

	dataFile, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	if len(dataFile) != 0 {

		err = json.Unmarshal(dataFile, &users)

		if err != nil {
			return err
		}
	}

	for _, u := range users {
		if u.ID == user.ID {
			return &errs.ErrAlreadyExists{Id: user.ID}
		}
	}

	users = append(users, user)

	dat, err := json.Marshal(users)

	if err != nil {
		return err
	}

	err = os.Truncate(s.Dir, 0)

	if err != nil {
		return err
	}

	file.Seek(0, io.SeekStart)
	file.Write(dat)

	return nil
}
