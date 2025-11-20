-- 初始化数据插入 -- 

-- 向 material_type 表插入数据
INSERT INTO `material_type` (`material_type_name`) VALUES
('Raw Material'),
('Finished Product'),
('Packaging'),
('Auxiliary Material');

-- 向 materials_info 表插入数据
INSERT INTO `materials_info` 
(`material_name`, `material_type_id`, `material_desc`, `material_status`, `material_purchased_at`, `material_location`) 
VALUES
('Steel Sheet', 1, 'High quality steel sheet for construction', 0, '2025-11-01 10:00:00', 'Warehouse A'),
('Aluminum Frame', 1, 'Aluminum frame for machinery', 2, '2025-10-25 14:30:00', 'Warehouse B'),
('Plastic Bottle', 3, 'Plastic bottle for packaging', 0, '2025-11-10 09:00:00', 'Warehouse C'),
('Finished Chair', 2, 'Wooden chair ready for retail', 1, '2025-11-15 16:00:00', 'Retail Store A'),
('Glass Bottle', 3, 'Glass bottle for beverages', 0, '2025-11-05 12:00:00', 'Warehouse A'),
('Wooden Panel', 1, 'Wooden panel for furniture production', 1, '2025-09-20 11:45:00', 'Warehouse B');

-- 确保外键检查重新启用
SET FOREIGN_KEY_CHECKS = 1;

