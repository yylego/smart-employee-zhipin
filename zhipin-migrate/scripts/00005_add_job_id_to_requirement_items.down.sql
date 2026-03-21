DROP INDEX IF EXISTS "idx_requirement_items_job_id";

ALTER TABLE "requirement_items" DROP COLUMN IF EXISTS "job_id";
