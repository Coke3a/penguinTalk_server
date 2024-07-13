CREATE TABLE "messages" (
    "msg_id" BIGSERIAL PRIMARY KEY,
    "convers_id" bigint NOT NULL,
    "user_id" bigint NOT NULL,
    "mt_id" bigint NOT NULL DEFAULT 0,
    "msg_text" varchar NOT NULL,
    "msg_audio" varchar,
    "msg_type" bigint NOT NULL,
    "msg_date" timestamptz NOT NULL DEFAULT now()
);