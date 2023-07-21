package auth

import (
	"github.com/HamelBarrer/online-calls-server/internal/core/user/user"
	"github.com/HamelBarrer/online-calls-server/internal/storage"
)

type Repository interface {
	GetUserByUsername(string) (*user.UserCreate, error)
}

type repository struct {
	s storage.SQL
}

func Newrepository(s storage.SQL) Repository {
	return &repository{s}
}

func (r *repository) GetUserByUsername(username string) (*user.UserCreate, error) {
	var u user.UserCreate

	query := `
		select
			u.user_id,
			u.username,
			u."password"
		from users.users u
		where u.username = $1;
	`

	if err := r.s.QueryRow(query, username).Scan(&u.UserId, &u.Username, &u.Password); err != nil {
		return nil, err
	}

	return &u, nil
}
