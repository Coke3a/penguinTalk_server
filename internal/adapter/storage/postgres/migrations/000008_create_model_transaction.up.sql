CREATE TABLE "model_transaction" (
    "mt_id" BIGSERIAL PRIMARY KEY,
    "request_prompt" TEXT NOT NULL,
    "response_data" TEXT NOT NULL,
    "transaction_date" timestamptz NOT NULL DEFAULT now()
);