package user

import userstate "github.com/HamelBarrer/online-calls-server/internal/core/user/user_state"

type User struct {
	UserId    int                 `json:"user_id,omitempty" required:"false"`
	Username  string              `json:"username,omitempty" required:"true"`
	UserState userstate.UserState `json:"user_state,omitempty" required:"false"`
}

type UserCreate struct {
	User
	Password        string `json:"password,omitempty" required:"true"`
	PasswordConfirm string `json:"password_confirm,omitempty" required:"true"`
}
