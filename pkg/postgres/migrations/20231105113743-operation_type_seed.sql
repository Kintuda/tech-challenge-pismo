-- +migrate Up
INSERT INTO public.operation_types
(id, operation_operator, description, created_at, deleted_at)
VALUES('db0c579e-fa2a-442c-bb21-48c1d65feeac', 'NEGATIVE', 'COMPRA A VISTA', now(), null);

INSERT INTO public.operation_types
(id, operation_operator, description, created_at, deleted_at)
VALUES('a767eb26-9efa-4e0f-908f-e184958f23e1', 'NEGATIVE', 'COMPRA PARCELADA', now(), null);

INSERT INTO public.operation_types
(id, operation_operator, description, created_at, deleted_at)
VALUES('80025b43-cd09-47fa-8dc3-fa2dc7f14d11', 'NEGATIVE', 'SAQUE', now(), null);

INSERT INTO public.operation_types
(id, operation_operator, description, created_at, deleted_at)
VALUES('c3822b88-6bcd-4a2b-9d7a-a2ac42ea364b', 'POSITIVE', 'PAGAMENTO', now(), null);
-- +migrate Down