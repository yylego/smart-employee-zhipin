ALTER TABLE "requirement_items" RENAME TO "match_items";

ALTER INDEX "idx_requirement_items_position_id" RENAME TO "idx_match_items_position_id";

ALTER INDEX "idx_requirement_items_deleted_at" RENAME TO "idx_match_items_deleted_at";
