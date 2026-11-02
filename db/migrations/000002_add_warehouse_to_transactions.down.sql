ALTER TABLE inbound_transactions
DROP FOREIGN KEY fk_inbound_warehouse,
DROP COLUMN warehouse_id;

ALTER TABLE outbound_transactions
DROP FOREIGN KEY fk_outbound_warehouse,
DROP COLUMN warehouse_id;
