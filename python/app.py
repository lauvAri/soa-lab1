from flask import Flask
from flask_cors import CORS
from config import Config
from models import db
from routes_borrows import borrows_bp

# 创建 Flask 应用
app = Flask(__name__)

# 加载配置
app.config.from_object(Config)

# 启用 CORS
CORS(app)

# 初始化数据库
db.init_app(app)

# 注册蓝图
app.register_blueprint(borrows_bp)


@app.route('/')
def index():
    """服务健康检查接口"""
    return {
        'code': 200,
        'message': 'Borrow Service is running',
        'service': 'borrow-service',
        'version': '1.0.0',
        'port': Config.PORT
    }


@app.route('/health')
def health():
    """健康检查接口"""
    return {
        'status': 'healthy',
        'service': 'borrow-service'
    }


# 创建数据库表
with app.app_context():
    db.create_all()
    print("数据库表创建成功!")


if __name__ == '__main__':
    print(f"借用记录服务启动中...")
    print(f"监听地址: {Config.HOST}:{Config.PORT}")
    print(f"数据库: {Config.SQLALCHEMY_DATABASE_URI}")
    print(f"用户服务: {Config.USER_SERVICE_BASE_URL}")
    print(f"物资服务: {Config.MATERIAL_SERVICE_BASE_URL}")
    app.run(host=Config.HOST, port=Config.PORT, debug=Config.DEBUG)
