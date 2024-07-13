CREATE TABLE "users" (
    "user_id" BIGSERIAL PRIMARY KEY,
    "username" varchar NOT NULL,
    "password" varchar NOT NULL,
    "email" varchar NOT NULL,
    "display_name" varchar NOT NULL,
    "users_rank" varchar NOT NULL, -- maybe fix it to int and create the table for user_rank
    "incorrect_login" bigint NOT NULL DEFAULT 0,
    "last_login" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

INSERT INTO "users" ("username", "password", "email", "display_name", "users_rank", "incorrect_login") 
VALUES ("test_username", "test_password", "test_email", "test_display_name", "test_user_rank");