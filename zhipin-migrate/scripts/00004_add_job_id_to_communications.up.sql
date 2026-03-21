ALTER TABLE "communications" ADD COLUMN IF NOT EXISTS "job_id" VARCHAR(28) NOT NULL DEFAULT '';

CREATE INDEX IF NOT EXISTS "idx_communications_job_id" ON "communications" ("job_id");
