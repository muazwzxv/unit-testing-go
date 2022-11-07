create table messages (
	id serial primary key not null,
	title varchar not null,
	body varchar not null,
	created_at timestamp not null
)
---- create above / drop below ----

drop table messages

