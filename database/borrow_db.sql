-- 借用记录服务数据库初始化脚本
-- 创建数据库
CREATE DATABASE IF NOT EXISTS borrow_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE borrow_db;

-- 借用记录表会由 SQLAlchemy 自动创建
-- 但如果需要手动创建，可以使用以下语句:

-- CREATE TABLE IF NOT EXISTS borrow_records (
--     id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '借用记录ID',
--     user_id BIGINT NOT NULL COMMENT '借用人ID',
--     material_id BIGINT NOT NULL COMMENT '物资ID',
--     quantity INT NOT NULL DEFAULT 1 COMMENT '借用数量',
--     status SMALLINT NOT NULL DEFAULT 0 COMMENT '借用状态: 0-借出中, 1-已归还, 2-已取消',
--     borrowed_at DATETIME NOT NULL COMMENT '借出时间',
--     due_at DATETIME NULL COMMENT '应归还时间',
--     returned_at DATETIME NULL COMMENT '实际归还时间',
--     remark VARCHAR(255) NULL COMMENT '备注信息',
--     created_at DATETIME NOT NULL COMMENT '创建时间',
--     updated_at DATETIME NOT NULL COMMENT '更新时间',
--     INDEX idx_user_id (user_id),
--     INDEX idx_material_id (material_id),
--     INDEX idx_status (status)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='借用记录表';

-- 插入测试数据（可选）
-- INSERT INTO borrow_records (user_id, material_id, quantity, status, borrowed_at, due_at, remark, created_at, updated_at)
-- VALUES 
--     (1, 1001, 1, 0, NOW(), DATE_ADD(NOW(), INTERVAL 7 DAY), '测试借用记录1', NOW(), NOW()),
--     (2, 1002, 1, 0, NOW(), DATE_ADD(NOW(), INTERVAL 7 DAY), '测试借用记录2', NOW(), NOW());
