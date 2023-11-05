-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "public"."accounts" (
    id UUID PRIMARY KEY default uuid_generate_v4(),
    document_number VARCHAR(14) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL default now(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE "public"."operation_types" (
    id UUID PRIMARY KEY default uuid_generate_v4(),
    description VARCHAR(256) NOT NULL,
    operation_operator VARCHAR(256) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL default now(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE "public"."transactions" (
    id UUID PRIMARY KEY default uuid_generate_v4(),
    account_id UUID NOT NULL,
    operation_type_id UUID NOT NULL,
    amount NUMERIC(14,2) NOT NULL,
    event_date TIMESTAMP WITH TIME ZONE NOT NULL default now(),
    CONSTRAINT fk_account FOREIGN KEY(account_id) REFERENCES accounts(id),
    CONSTRAINT fk_operation_type FOREIGN KEY(operation_type_id) REFERENCES operation_types(id)
);
-- +migrate Down
DROP TABLE "public"."accounts";
DROP TABLE  "public"."operation_types";
DROP TABLE "public"."transactions";