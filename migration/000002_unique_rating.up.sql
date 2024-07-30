ALTER TABLE user_rating
    ADD CONSTRAINT tg_id_unique UNIQUE (tg_id);