package channelstates

type ChannelState struct {
	ChannelStateId int    `json:"channel_state_id,omitempty" required:"false"`
	Name           string `json:"name,omitempty" required:"true"`
}
