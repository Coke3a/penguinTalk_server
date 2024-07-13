CREATE TABLE "prompt_language" (
    "prompt_lang_id" BIGSERIAL PRIMARY KEY,
    "language" varchar NOT NULL,
    "status" bigint NOT NULL
);  

INSERT INTO "prompt_language" ("language", "status") VALUES ('English', 1);