CREATE TABLE Authentication(
    id              UUID         CONSTRAINT authentication_pk primary key,
    token           varchar(255) NOT NULL,
    updated_at      timestamp    NOT NULL default CURRENT_TIMESTAMP,
    expiration_time integer      NOT NULL
);