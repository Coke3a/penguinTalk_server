CREATE TABLE "users" (
    "user_id" BIGSERIAL PRIMARY KEY,
    "username" varchar NOT NULL,
    "password" varchar NOT NULL,
    "email" varchar NOT NULL,
    "users_rank" varchar NOT NULL, -- maybe fix it to int and create the table for user_rank
    "incorrect_login" bigint NOT NULL DEFAULT 0,
    "last_login" timestamptz,
    "created_date" timestamptz NOT NULL DEFAULT (now())
);