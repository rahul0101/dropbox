CREATE TABLE IF NOT EXISTS file_meta_data
(
    id               character varying(40) NOT NULL PRIMARY KEY,
    name             character varying(40) NOT NULL,
    creator_id       character varying(40) NOT NULL,
    created_at_epoch bigint                NOT NULL
);


CREATE INDEX IF NOT EXISTS file_meta_data_creator_id_idx ON file_meta_data (creator_id);