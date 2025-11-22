from flask_sqlalchemy import SQLAlchemy
from datetime import datetime

db = SQLAlchemy()


class BorrowRecord(db.Model):
    """借用记录模型"""
    
    __tablename__ = 'borrow_records'
    
    # 主键
    id = db.Column(db.BigInteger, primary_key=True, autoincrement=True, comment='借用记录ID')
    
    # 外键关联
    user_id = db.Column(db.BigInteger, nullable=False, index=True, comment='借用人ID')
    material_id = db.Column(db.BigInteger, nullable=False, index=True, comment='物资ID')
    
    # 借用信息
    quantity = db.Column(db.Integer, nullable=False, default=1, comment='借用数量')
    status = db.Column(db.SmallInteger, nullable=False, default=0, index=True, comment='借用状态: 0-借出中, 1-已归还, 2-已取消')
    
    # 时间字段
    borrowed_at = db.Column(db.DateTime, nullable=False, default=datetime.utcnow, comment='借出时间')
    due_at = db.Column(db.DateTime, nullable=True, comment='应归还时间')
    returned_at = db.Column(db.DateTime, nullable=True, comment='实际归还时间')
    
    # 备注
    remark = db.Column(db.String(255), nullable=True, comment='备注信息')
    
    # 审计字段
    created_at = db.Column(db.DateTime, nullable=False, default=datetime.utcnow, comment='创建时间')
    updated_at = db.Column(db.DateTime, nullable=False, default=datetime.utcnow, onupdate=datetime.utcnow, comment='更新时间')
    
    # 索引
    __table_args__ = (
        db.Index('idx_user_id', 'user_id'),
        db.Index('idx_material_id', 'material_id'),
        db.Index('idx_status', 'status'),
    )
    
    # 状态枚举
    STATUS_BORROWED = 0  # 借出中
    STATUS_RETURNED = 1  # 已归还
    STATUS_CANCELLED = 2  # 已取消
    
    STATUS_TEXT = {
        0: '借出中',
        1: '已归还',
        2: '已取消'
    }
    
    def to_dict(self, include_user=False, include_material=False, user_data=None, material_data=None):
        """
        将模型转换为字典
        
        Args:
            include_user: 是否包含用户信息
            include_material: 是否包含物资信息
            user_data: 用户数据字典
            material_data: 物资数据字典
        """
        result = {
            'id': self.id,
            'userId': self.user_id,
            'materialId': self.material_id,
            'quantity': self.quantity,
            'status': self.status,
            'statusText': self.STATUS_TEXT.get(self.status, '未知'),
            'borrowedAt': self.borrowed_at.isoformat() if self.borrowed_at else None,
            'dueAt': self.due_at.isoformat() if self.due_at else None,
            'returnedAt': self.returned_at.isoformat() if self.returned_at else None,
            'remark': self.remark
        }
        
        # 如果需要包含用户信息
        if include_user and user_data:
            result['user'] = user_data
        
        # 如果需要包含物资信息
        if include_material and material_data:
            result['material'] = material_data
        
        return result
    
    def __repr__(self):
        return f'<BorrowRecord {self.id}: User {self.user_id} borrowed Material {self.material_id}>'
