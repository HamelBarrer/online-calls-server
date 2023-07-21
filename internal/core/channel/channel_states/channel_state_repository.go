package channelstates

import "github.com/HamelBarrer/online-calls-server/internal/storage"

type Repository interface {
	GetById(int) (*ChannelState, error)
	GetAll() (*[]ChannelState, error)
	Create(ChannelState) (*ChannelState, error)
}

type repository struct {
	s storage.SQL
}

func Newrepository(s storage.SQL) Repository {
	return &repository{s}
}

func (r *repository) GetById(channelStateId int) (*ChannelState, error) {
	var cs ChannelState

	query := `
		select
			cs.channel_state_id,
			cs."name"
		from channels.channel_states cs
		where cs.channel_state_id = $1;
	`

	if err := r.s.QueryRow(query, channelStateId).Scan(&cs.ChannelStateId, &cs.Name); err != nil {
		return nil, err
	}

	return &cs, nil
}

func (r *repository) GetAll() (*[]ChannelState, error) {
	var css []ChannelState

	query := `
		select
			cs.channel_state_id,
			cs."name"
		from channels.channel_states cs;
	`

	rows, err := r.s.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var cs ChannelState

		if err := rows.Scan(&cs.ChannelStateId, &cs.Name); err != nil {
			return &[]ChannelState{}, err
		}

		css = append(css, cs)
	}

	return &css, nil
}

func (r *repository) Create(channelState ChannelState) (*ChannelState, error) {
	lastInsertId := 0

	query := `
		insert into channels.channel_states (name)
		values ($1)
		returning channel_state_id;
	`
	if err := r.s.QueryRow(query, channelState.Name).Scan(&lastInsertId); err != nil {
		return nil, err
	}

	return r.GetById(lastInsertId)
}
