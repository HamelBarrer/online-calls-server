package channels

import (
	channelstates "github.com/HamelBarrer/online-calls-server/internal/core/channel/channel_states"
	"github.com/HamelBarrer/online-calls-server/internal/core/user/user"
)

type Channel struct {
	ChannelId      int                        `json:"channel_id,omitempty" required:"false"`
	Name           string                     `json:"name,omitempty" required:"true"`
	ChannelState   channelstates.ChannelState `json:"channel_state,omitempty" required:"false"`
	ChannelStateId int                        `json:"channel_state_id,omitempty" required:"true"`
	CreatedUser    user.User                  `json:"created_user,omitempty" required:"false"`
	CreatedUserId  int                        `json:"created_user_id,omitempty" required:"true"`
}
