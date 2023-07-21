package channels

import "github.com/HamelBarrer/online-calls-server/internal/storage"

type Repository interface {
	GetById(int) (*Channel, error)
	GetAll() (*[]Channel, error)
	Create(Channel) (*Channel, error)
}

type repository struct {
	s storage.SQL
}

func Newrepository(s storage.SQL) Repository {
	return &repository{s}
}

func (r *repository) GetById(channelId int) (*Channel, error) {
	var c Channel

	query := `
		select
			c.channel_id,
			c."name",
			cs.channel_state_id,
			cs."name",
			u.user_id,
			u.username
		from channels.channels c
			join channels.channel_states cs using (channel_state_id)
			join users.users u on u.user_id = c.created_user_id
		where c.channel_id = $1;
	`

	err := r.s.QueryRow(query, channelId).Scan(
		&c.ChannelId,
		&c.Name,
		&c.ChannelState.ChannelStateId,
		&c.ChannelState.Name,
		&c.CreatedUser.UserId,
		&c.CreatedUser.Username,
	)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *repository) GetAll() (*[]Channel, error) {
	var cs []Channel

	query := `
		select
			c.channel_id,
			c."name",
			cs.channel_state_id,
			cs."name",
			u.user_id,
			u.username
		from channels.channels c
			join channels.channel_states cs using (channel_state_id)
			join users.users u on u.user_id = c.created_user_id;
	`

	rows, err := r.s.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c Channel

		err := rows.Scan(
			&c.ChannelId,
			&c.Name,
			&c.ChannelState.ChannelStateId,
			&c.ChannelState.Name,
			&c.CreatedUser.UserId,
			&c.CreatedUser.Username,
		)
		if err != nil {
			return nil, err
		}

		cs = append(cs, c)
	}

	return &cs, nil
}

func (r *repository) Create(c Channel) (*Channel, error) {
	insertedId := 0

	query := `
		insert into channels.channels (name, channel_state_id, created_user_id)
		values ($1, $2, $3)
		returning channel_id;
	`

	if err := r.s.QueryRow(query, c.Name, c.ChannelStateId, c.CreatedUserId).Scan(&insertedId); err != nil {
		return nil, err
	}

	return r.GetById(insertedId)
}
