package channelmessages

import (
	"github.com/HamelBarrer/online-calls-server/internal/core/channel/channels"
	"github.com/HamelBarrer/online-calls-server/internal/core/user/user"
)

type ChannelMessage struct {
	ChannelMessageId int              `json:"channel_message_id,omitempty" required:"false"`
	Channel          channels.Channel `json:"channel,omitempty" required:"false"`
	ChannelId        int              `json:"channel_id,omitempty" required:"true"`
	User             user.User        `json:"user,omitempty" required:"false"`
	UserId           int              `json:"user_id,omitempty" required:"true"`
	Message          string           `json:"message,omitempty" required:"true"`
}
