import os
from dotenv import load_dotenv

# 加载 .env 文件
load_dotenv()


class Config:
    """Flask 应用配置类"""
    
    # Flask 配置
    SECRET_KEY = os.getenv('SECRET_KEY', 'dev-secret-key-change-in-production')
    DEBUG = os.getenv('DEBUG', 'True').lower() == 'true'
    
    # 服务器配置
    HOST = os.getenv('HOST', '0.0.0.0')
    PORT = int(os.getenv('PORT', '8081'))
    
    # 数据库配置
    DB_HOST = os.getenv('DB_HOST', 'localhost')
    DB_PORT = int(os.getenv('DB_PORT', '3306'))
    DB_USER = os.getenv('DB_USER', 'root')
    DB_PASSWORD = os.getenv('DB_PASSWORD', 'root')
    DB_NAME = os.getenv('DB_NAME', 'borrow_db')
    
    # SQLAlchemy 配置
    SQLALCHEMY_DATABASE_URI = (
        f"mysql+pymysql://{DB_USER}:{DB_PASSWORD}@{DB_HOST}:{DB_PORT}/{DB_NAME}"
        "?charset=utf8mb4"
    )
    SQLALCHEMY_TRACK_MODIFICATIONS = False
    SQLALCHEMY_ECHO = DEBUG
    
    # 其他服务地址
    USER_SERVICE_BASE_URL = os.getenv('USER_SERVICE_BASE_URL', 'http://localhost:8083')
    MATERIAL_SERVICE_BASE_URL = os.getenv('MATERIAL_SERVICE_BASE_URL', 'http://localhost:8082')
    
    # 分页配置
    DEFAULT_PAGE_SIZE = 10
    MAX_PAGE_SIZE = 100
