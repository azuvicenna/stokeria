ALTER TABLE inbound_transactions
ADD COLUMN warehouse_id CHAR(36) NOT NULL AFTER supplier_id,
ADD CONSTRAINT fk_inbound_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE RESTRICT;

ALTER TABLE outbound_transactions
ADD COLUMN warehouse_id CHAR(36) NOT NULL AFTER reference_code,
ADD CONSTRAINT fk_outbound_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE RESTRICT;
