from flask import Flask
from flask_cors import CORS
from flasgger import Swagger
from config import Config
from models import db
from routes_borrows import borrows_bp

# 创建 Flask 应用
app = Flask(__name__)

# 加载配置
app.config.from_object(Config)

# 启用 CORS
CORS(app)

# 初始化 Swagger 文档
swagger_template = {
    "swagger": "2.0",
    "info": {
        "title": "Borrow Service API",
        "description": "借用记录服务接口文档，提供借用记录的创建、查询、更新、归还、删除等能力。",
        "contact": {
            "name": "Borrow Service Team"
        },
        "version": "1.0.0"
    },
    "basePath": "/",
    "schemes": ["http", "https"],
    "tags": [
        {"name": "Health", "description": "服务健康检查"},
        {"name": "Borrows", "description": "借用记录管理接口"}
    ],
    "definitions": {
        "BorrowRecord": {
            "type": "object",
            "properties": {
                "id": {"type": "integer", "format": "int64"},
                "userId": {"type": "integer", "format": "int64"},
                "materialId": {"type": "integer", "format": "int64"},
                "quantity": {"type": "integer"},
                "status": {"type": "integer", "description": "0-借出中, 1-已归还, 2-已取消"},
                "statusText": {"type": "string"},
                "borrowedAt": {"type": "string", "format": "date-time"},
                "dueAt": {"type": "string", "format": "date-time"},
                "returnedAt": {"type": "string", "format": "date-time"},
                "remark": {"type": "string"},
                "user": {"type": "object"},
                "material": {"type": "object"}
            }
        },
        "BorrowCreateRequest": {
            "type": "object",
            "required": ["userId", "materialId"],
            "properties": {
                "userId": {"type": "integer", "format": "int64"},
                "materialId": {"type": "integer", "format": "int64"},
                "quantity": {"type": "integer", "default": 1},
                "dueAt": {"type": "string", "format": "date-time"},
                "remark": {"type": "string"}
            }
        },
        "BorrowUpdateRequest": {
            "type": "object",
            "properties": {
                "status": {"type": "integer", "enum": [0, 1, 2]},
                "returnedAt": {"type": "string", "format": "date-time"},
                "dueAt": {"type": "string", "format": "date-time"},
                "remark": {"type": "string"}
            }
        },
        "BorrowReturnRequest": {
            "type": "object",
            "properties": {
                "returnedAt": {"type": "string", "format": "date-time"},
                "remark": {"type": "string"}
            }
        },
        "PagedBorrowResult": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {"$ref": "#/definitions/BorrowRecord"}
                },
                "page": {"type": "integer"},
                "pageSize": {"type": "integer"},
                "total": {"type": "integer"}
            }
        },
        "BaseResponse": {
            "type": "object",
            "properties": {
                "code": {"type": "integer"},
                "message": {"type": "string"},
                "timestamp": {"type": "string", "format": "date-time"}
            }
        },
        "BorrowRecordResponse": {
            "allOf": [
                {"$ref": "#/definitions/BaseResponse"},
                {
                    "type": "object",
                    "properties": {
                        "data": {"$ref": "#/definitions/BorrowRecord"}
                    }
                }
            ]
        },
        "BorrowListResponse": {
            "allOf": [
                {"$ref": "#/definitions/BaseResponse"},
                {
                    "type": "object",
                    "properties": {
                        "data": {"$ref": "#/definitions/PagedBorrowResult"}
                    }
                }
            ]
        },
        "SimpleResponse": {
            "allOf": [
                {"$ref": "#/definitions/BaseResponse"},
                {
                    "type": "object",
                    "properties": {
                        "data": {"type": ["object", "null"]}
                    }
                }
            ]
        }
    }
}

swagger = Swagger(app, template=swagger_template)

# 初始化数据库
db.init_app(app)

# 注册蓝图
app.register_blueprint(borrows_bp)


@app.route('/')
def index():
    """
    服务健康检查接口
    ---
    tags:
      - Health
    responses:
      200:
        description: 服务健康状态
        schema:
          type: object
          properties:
            code:
              type: integer
              example: 200
            message:
              type: string
            service:
              type: string
            version:
              type: string
            port:
              type: integer
    """
    return {
        'code': 200,
        'message': 'Borrow Service is running',
        'service': 'borrow-service',
        'version': '1.0.0',
        'port': Config.PORT
    }


@app.route('/health')
def health():
    """
    健康检查接口
    ---
    tags:
      - Health
    responses:
      200:
        description: 服务健康状态
        schema:
          type: object
          properties:
            status:
              type: string
            service:
              type: string
          required:
            - status
            - service
    """
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
