CREATE DATABASE vote_account;
use vote_account;
CREATE TABLE users (
    id bigint not null,
    guid varchar(38) not null,
    metadata text not null,
    hash text not null
);
CREATE TABLE voters (
    userId bigint not null,
    nonce bigint not null,
    prevHash text not null,
    hash text not null
);
