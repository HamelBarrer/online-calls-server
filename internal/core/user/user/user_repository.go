package user

import (
	"github.com/HamelBarrer/online-calls-server/internal/storage"
	"github.com/HamelBarrer/online-calls-server/internal/utils"
)

type Repository interface {
	GetById(int) (*User, error)
	GetAll() (*[]User, error)
	Create(UserCreate) (*User, error)
}

type repository struct {
	p storage.SQL
}

func Newrepository(p storage.SQL) Repository {
	return &repository{p}
}

func (r *repository) GetById(userId int) (*User, error) {
	u := User{}

	query := `
		select
			u.user_id,
			u.username,
			us.user_state_id,
			us."name" name_state
		from users.users u
			join users.user_states us using(user_state_id)
		where u.user_id = $1;
	`

	err := r.p.QueryRow(query, userId).Scan(&u.UserId, &u.Username, &u.UserState.UserStateId, &u.UserState.Name)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *repository) GetAll() (*[]User, error) {
	query := `
		select
			u.user_id,
			u.username,
			us.user_state_id,
			us."name" name_state
		from users.users u
			join users.user_states us using(user_state_id);
	`

	rows, _ := r.p.Query(query)

	u := User{}
	us := []User{}

	for rows.Next() {

		if err := rows.Scan(&u.UserId, &u.Username, &u.UserState.UserStateId, &u.UserState.Name); err != nil {
			return &[]User{}, err
		}

		us = append(us, u)
	}

	return &us, nil
}

func (r *repository) Create(uc UserCreate) (*User, error) {
	p, err := utils.CreationHash(uc.Password)
	if err != nil {
		return nil, err
	}

	query := `
		insert into users.users (username, password)
		values ($1, $2)
		returning user_id;
	`

	userIdRegister := 0

	err = r.p.QueryRow(query, uc.Username, p).Scan(&userIdRegister)
	if err != nil {
		return nil, err
	}

	return r.GetById(userIdRegister)
}
