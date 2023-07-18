create table if not exists users.user_states (
	user_state_id bigserial not null primary key,
	name varchar not null unique,
	is_active bool not null default true,
	created_at timestamp not null default now(),
	updated_at timestamp default now()
);

create table if not exists users.users (
	user_id bigserial not null primary key,
	username varchar not null unique,
	password varchar not null,
	user_state_id int not null references users.user_states on delete cascade on update cascade default 1,
	created_at timestamp not null default now(),
	updated_at timestamp default now()
);

create table if not exists channels.channel_states (
	channel_state_id bigserial not null primary key,
	name varchar not null unique,
	is_active bool not null default true,
	created_at timestamp not null default now(),
	updated_at timestamp default now()
);

create table if not exists channels.channels (
	channel_id bigserial not null primary key,
	name varchar not null unique,
	channel_state_id int not null references channels.channel_states on delete cascade on update cascade,
	created_user_id int not null references users.users(user_id) on delete cascade on update cascade,
	created_at timestamp not null default now(),
	updated_at timestamp default now()
);

create table if not exists channels.channel_users (
	channel_user_id bigserial not null primary key,
	channel_id int not null references channels.channels on delete cascade on update cascade,
	user_id int not null references users.users on delete cascade on update cascade,
	created_at timestamp not null default now(),
	updated_at timestamp default now()
);

create table if not exists channels.channel_messages (
	channel_message_id bigserial not null primary key,
	channel_id int not null references channels.channels on delete cascade on update cascade,
	user_id int not null references users.users on delete cascade on update cascade,
	message varchar not null,
	created_at timestamp not null default now(),
	updated_at timestamp default now()
);

create table if not exists channels.channel_roles (
	channel_role_id bigserial not null primary key,
	channel_id int not null references channels.channels on delete cascade on update cascade,
	name varchar not null,
	is_active bool not null default true,
	created_at timestamp not null default now(),
	updated_at timestamp default now()
);

create table if not exists channels.channel_role_users (
	channel_role_user_id bigserial not null primary key,
	channel_role_id int not null references channels.channel_roles on delete cascade on update cascade,
	user_id int not null references users.users on delete cascade on update cascade,
	assigned_user_id int not null references users.users (user_id) on delete cascade on update cascade,
	created_at timestamp not null default now(),
	updated_at timestamp default now()
);