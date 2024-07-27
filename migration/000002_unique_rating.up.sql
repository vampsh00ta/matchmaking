ALTER TABLE user_matchmaking
    ADD CONSTRAINT tg_id_unique UNIQUE (tg_id);