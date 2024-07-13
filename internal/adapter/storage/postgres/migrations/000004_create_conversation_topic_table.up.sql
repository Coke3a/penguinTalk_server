CREATE TABLE "conversation_topic" (
    "convers_topic_id" BIGSERIAL PRIMARY KEY,
    "convers_topic_name" varchar NOT NULL,
    "convers_topic_description" varchar
);

INSERT INTO "conversation_topic" ("convers_topic_name", "convers_topic_description") 
VALUES ('General Discussion', 'Topics for general conversation');