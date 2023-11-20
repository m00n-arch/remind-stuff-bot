package db

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type status struct {
	userID string
	state  string
}

type UsersDB struct {
	f     *os.File
	users []status
}

func NewUserDB(path string) (*UsersDB, error) {
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	u := &UsersDB{
		f:     f,
		users: nil,
	}

	return u, u.load()
}

const (
	StartState = "start"
)

func (u *UsersDB) GetState(userIDToFind string) (string, error) {
	for i := range u.users {
		if u.users[i].userID == userIDToFind {
			return u.users[i].state, nil
		}
	}
	return StartState, u.UpdateState(userIDToFind, StartState)
}

func (u *UsersDB) UpdateState(userIDToFind string, newState string) error {
	for i := range u.users {
		if u.users[i].userID == userIDToFind {
			u.users[i].state = newState
			return u.save()
		}
	}
	u.users = append(u.users, status{
		userID: userIDToFind,
		state:  newState,
	})

	return u.save()
}

func (u *UsersDB) save() error {
	_, err := u.f.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("can't reset offset: %w", err)
	}

	err = u.f.Truncate(0)
	if err != nil {
		return fmt.Errorf("can't truncate the file: %w", err)
	}

	for i := range u.users {
		_, err := u.f.Write([]byte(fmt.Sprintf("%s,%s\n", u.users[i].userID, u.users[i].state)))
		if err != nil {
			return err
		}
	}

	return u.f.Sync()
}

func (u *UsersDB) load() error {
	u.users = nil

	_, err := u.f.Seek(0, 0)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(u.f)
	if err != nil {
		return err
	}
	str := string(data)

	rows := strings.Split(str, "\n")

	for i := range rows {
		if i == len(rows)-1 {
			continue
		}

		cells := strings.Split(rows[i], ",")
		if len(cells) != 2 {
			return fmt.Errorf("invalid cells count")
		}
		s := status{userID: cells[0], state: cells[1]}
		u.users = append(u.users, s)
	}

	return nil
}
