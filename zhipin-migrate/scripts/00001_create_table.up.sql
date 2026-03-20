CREATE TABLE "positions" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"job_id" varchar(28) NOT NULL,"title" varchar(256) NOT NULL,"company" varchar(256) NOT NULL,"salary_range" varchar(64) NOT NULL,"salary_min" integer,"salary_max" integer,"city" varchar(64) NOT NULL,"link" varchar(512),"recruiter" varchar(128),"enc_boss_id" varchar(64),"is_hunter" boolean,"status" varchar(32),"skip_reason" text,"match_rate" integer,"notes" text,"last_comm_at" bigint,"last_comm_dir" integer,"last_resume" varchar(256),"duties" text,"requirements" text,PRIMARY KEY ("id"));

CREATE UNIQUE INDEX IF NOT EXISTS "idx_positions_job_id" ON "positions" ("job_id");

CREATE INDEX IF NOT EXISTS "idx_positions_deleted_at" ON "positions" ("deleted_at");

CREATE TABLE "match_items" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"position_id" bigint NOT NULL,"requirement" text NOT NULL,"match_status" varchar(32) NOT NULL,"resume_point" text,"remark" varchar(512),"sort_index" integer,PRIMARY KEY ("id"));

CREATE INDEX IF NOT EXISTS "idx_match_items_position_id" ON "match_items" ("position_id");

CREATE INDEX IF NOT EXISTS "idx_match_items_deleted_at" ON "match_items" ("deleted_at");

CREATE TABLE "communications" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"position_id" bigint NOT NULL,"event_type" varchar(32) NOT NULL,"event_time" timestamptz NOT NULL,"content" text,"direction" integer,PRIMARY KEY ("id"));

CREATE INDEX IF NOT EXISTS "idx_communications_position_id" ON "communications" ("position_id");

CREATE INDEX IF NOT EXISTS "idx_communications_deleted_at" ON "communications" ("deleted_at");

CREATE TABLE "blacklist" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"company" varchar(256) NOT NULL,"reason" text,PRIMARY KEY ("id"));

CREATE UNIQUE INDEX IF NOT EXISTS "idx_blacklist_company" ON "blacklist" ("company");

CREATE INDEX IF NOT EXISTS "idx_blacklist_deleted_at" ON "blacklist" ("deleted_at");
