"""
数据库初始化脚本
用于单独创建数据库表
"""
from app import app, db

if __name__ == '__main__':
    with app.app_context():
        # 删除所有表（谨慎使用！）
        # db.drop_all()
        
        # 创建所有表
        db.create_all()
        print("✓ 数据库表创建成功!")
        
        # 显示创建的表
        from sqlalchemy import inspect
        inspector = inspect(db.engine)
        tables = inspector.get_table_names()
        print(f"✓ 已创建的表: {', '.join(tables)}")
