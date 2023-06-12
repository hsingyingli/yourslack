CREATE TABLE "slack_workspaces" (
  "id" bigint PRIMARY KEY,
  "title" varchar NOT NULL,
  "token" varchar
);


CREATE TABLE "slack_channels" (
  "id" bigint PRIMARY KEY,
  "workspace_id" bigint NOT NULL,
  "name" varchar NOT NULL
);

ALTER TABLE "slack_channels"
ADD FOREIGN KEY ("workspace_id")
REFERENCES "slack_workspaces" ("id")
ON DELETE CASCADE;
