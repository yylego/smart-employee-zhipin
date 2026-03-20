DROP INDEX IF EXISTS "idx_blacklist_deleted_at";
DROP INDEX IF EXISTS "idx_blacklist_company";
DROP TABLE IF EXISTS "blacklist";

DROP INDEX IF EXISTS "idx_communications_deleted_at";
DROP INDEX IF EXISTS "idx_communications_position_id";
DROP TABLE IF EXISTS "communications";

DROP INDEX IF EXISTS "idx_match_items_deleted_at";
DROP INDEX IF EXISTS "idx_match_items_position_id";
DROP TABLE IF EXISTS "match_items";

DROP INDEX IF EXISTS "idx_positions_deleted_at";
DROP INDEX IF EXISTS "idx_positions_job_id";
DROP TABLE IF EXISTS "positions";
