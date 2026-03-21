DROP INDEX IF EXISTS "idx_communications_job_id";

ALTER TABLE "communications" DROP COLUMN IF EXISTS "job_id";
