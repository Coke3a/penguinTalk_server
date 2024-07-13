CREATE TABLE "role" (
    "role_id" BIGSERIAL PRIMARY KEY,
    "role" varchar NOT NULL
);


INSERT INTO "role" ("role") VALUES ('User');
INSERT INTO "role" ("role") VALUES ('Assistant');
INSERT INTO "role" ("role") VALUES ('Best Friend');