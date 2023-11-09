-- +migrate Up
ALTER TABLE "public"."transactions"
    ADD COLUMN balance numeric(14, 2) NOT NULL DEFAULT '0';

-- +migrate Down

ALTER TABLE "public"."transactions"
    DROP COLUMN balance;