CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "u_id" bigint NOT NULL,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "expired_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT(now())
);

ALTER TABLE "sessions" 
ADD FOREIGN KEY ("u_id") 
REFERENCES "users" ("id")
ON DELETE CASCADE;
