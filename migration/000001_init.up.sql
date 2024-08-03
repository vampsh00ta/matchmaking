begin;

create table user_rating (
    id serial primary key ,
    rating integer,
    tg_id bigint
);

commit;