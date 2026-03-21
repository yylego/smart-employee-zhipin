ALTER TABLE "match_items" RENAME TO "requirement_items";

ALTER INDEX "idx_match_items_position_id" RENAME TO "idx_requirement_items_position_id";

ALTER INDEX "idx_match_items_deleted_at" RENAME TO "idx_requirement_items_deleted_at";
