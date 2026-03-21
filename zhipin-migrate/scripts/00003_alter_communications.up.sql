ALTER TABLE "communications" DROP COLUMN IF EXISTS "event_type";
ALTER TABLE "communications" DROP COLUMN IF EXISTS "event_time";

ALTER TABLE "communications" ADD COLUMN IF NOT EXISTS "timestamp" BIGINT NOT NULL DEFAULT 0;
ALTER TABLE "communications" ADD COLUMN IF NOT EXISTS "is_resume" BOOLEAN DEFAULT FALSE;
ALTER TABLE "communications" ADD COLUMN IF NOT EXISTS "resume_version" VARCHAR(256);

ALTER TABLE "communications" ALTER COLUMN "content" SET NOT NULL;
ALTER TABLE "communications" ALTER COLUMN "direction" SET NOT NULL;
