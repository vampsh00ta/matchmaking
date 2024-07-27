begin;

create table user_matchmaking (
    id serial primary key ,
    matchmaking integer,
    tg_id integer
);
commit;