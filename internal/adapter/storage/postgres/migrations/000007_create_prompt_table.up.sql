CREATE TABLE "prompt" (
    "prompt_id" BIGSERIAL PRIMARY KEY,
    "convers_topic_id" bigint NOT NULL,
    "prompt_lang_id" bigint NOT NULL,
    "prompt" varchar NOT NULL,
    "prompt_2" varchar,
    "ai_role" bigint NOT NULL,
    "user_role" bigint NOT NULL,
    FOREIGN KEY ("convers_topic_id") REFERENCES "conversation_topic" ("convers_topic_id"),
    FOREIGN KEY ("prompt_lang_id") REFERENCES "prompt_language" ("prompt_lang_id"),
    FOREIGN KEY ("ai_role") REFERENCES "role" ("role_id"),
    FOREIGN KEY ("user_role") REFERENCES "role" ("role_id")
);

INSERT INTO "prompt" ("convers_topic_id", "prompt_lang_id", "prompt", "prompt_2", "ai_role", "user_role") 
VALUES (1, 1, 'You’re my best friend. Start a conversation with me by asking a question or talking about something interesting. If I send the message “‘empty’”, you should initiate the conversation. Ensure the conversation remains fully in character and natural, without breaking the role-play by mentioning the chat setting. Keep the responses short and natural, like how friends talk to each other.', 'Example Your task is to engage in a friendly conversation.', 1, 2);