ALTER TABLE "requirement_items" ADD COLUMN IF NOT EXISTS "job_id" VARCHAR(28) NOT NULL DEFAULT '';

CREATE INDEX IF NOT EXISTS "idx_requirement_items_job_id" ON "requirement_items" ("job_id");
