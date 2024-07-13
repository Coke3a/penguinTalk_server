CREATE TABLE "conversation" (
    "convers_id" BIGSERIAL PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "prompt_id" bigint NOT NULL,
    "convers_start" timestamptz NOT NULL DEFAULT now(),
    "convers_end" timestamptz
);