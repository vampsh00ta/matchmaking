begin;

create table user_rating (
    id serial primary key ,
    matchmaking integer,
    tg_id bigint
);

commit;