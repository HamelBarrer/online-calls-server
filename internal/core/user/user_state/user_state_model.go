package userstate

type UserState struct {
	UserStateId int    `json:"user_state_id,omitempty" required:"false"`
	Name        string `json:"name,omitempty" required:"true"`
}
