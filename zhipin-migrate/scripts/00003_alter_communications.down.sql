ALTER TABLE "communications" DROP COLUMN IF EXISTS "timestamp";
ALTER TABLE "communications" DROP COLUMN IF EXISTS "is_resume";
ALTER TABLE "communications" DROP COLUMN IF EXISTS "resume_version";

ALTER TABLE "communications" ADD COLUMN IF NOT EXISTS "event_type" VARCHAR(32) NOT NULL DEFAULT '';
ALTER TABLE "communications" ADD COLUMN IF NOT EXISTS "event_time" TIMESTAMPTZ NOT NULL DEFAULT NOW();

ALTER TABLE "communications" ALTER COLUMN "content" DROP NOT NULL;
ALTER TABLE "communications" ALTER COLUMN "direction" DROP NOT NULL;
