package channelmessages

import "github.com/HamelBarrer/online-calls-server/internal/storage"

type Repository interface {
	GetById(int) (*ChannelMessage, error)
	GetAll() (*[]ChannelMessage, error)
	Create(ChannelMessage) (*ChannelMessage, error)
}

type repository struct {
	s storage.SQL
}

func Newrepository(s storage.SQL) Repository {
	return &repository{s}
}

func (r *repository) GetById(channelMessageId int) (*ChannelMessage, error) {
	var cm ChannelMessage

	query := `
		select
			cm.channel_message_id,
			c.channel_id,
			c."name",
			u.user_id,
			u.username,
			cm.message
		from channels.channel_messages cm
			join channels.channels c using (channel_id)
			join users.users u using (user_id)
		where c.channel_id = $1;
	`

	err := r.s.QueryRow(query, channelMessageId).Scan(
		&cm.ChannelMessageId,
		&cm.Channel.ChannelId,
		&cm.Channel.Name,
		&cm.User.UserId,
		&cm.User.Username,
		&cm.Message,
	)
	if err != nil {
		return nil, err
	}

	return &cm, nil
}

func (r *repository) GetAll() (*[]ChannelMessage, error) {
	cms := []ChannelMessage{}

	query := `
		select
			cm.channel_message_id,
			c.channel_id,
			c."name",
			u.user_id,
			u.username,
			cm.message
		from channels.channel_messages cm
			join channels.channels c using (channel_id)
			join users.users u using (user_id)
		where c.channel_id = $1;
	`

	rows, err := r.s.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		cm := ChannelMessage{}

		err := rows.Scan(
			&cm.ChannelMessageId,
			&cm.Channel.ChannelId,
			&cm.Channel.Name,
			&cm.User.UserId,
			&cm.User.Username,
			&cm.Message,
		)
		if err != nil {
			return nil, err
		}

		cms = append(cms, cm)
	}

	return &cms, nil
}

func (r *repository) Create(channelMessage ChannelMessage) (*ChannelMessage, error) {
	registerId := 0

	query := `
		insert into channels.channel_messages (channel_id, user_id, message)
		values ($1, $2, $3)
		returning channel_message_id;
	`

	err := r.s.QueryRow(query, channelMessage.ChannelId, channelMessage.UserId, channelMessage.Message).Scan(&registerId)
	if err != nil {
		return nil, err
	}

	return r.GetById(registerId)
}
