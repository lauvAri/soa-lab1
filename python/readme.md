# 借用记录服务 (Borrow Service)

Python + Flask 实现的借用记录管理服务,运行在端口 **8081**。

## 服务概述

- **服务名称**: 借用记录服务 (Borrow Service)
- **技术栈**: Python 3 + Flask + SQLAlchemy + MySQL
- **运行端口**: 8081
- **职责**: 
  - 提供借用记录的增删改查 REST API
  - 实现借出、归还、查询借用历史等业务逻辑
  - 与人员管理服务 (Java/8083) 和物资管理服务 (Go/8082) 协作

## 前置要求

- Python >= 3.9
- MySQL >= 5.7
- 人员管理服务 (8083) 正常运行
- 物资管理服务 (8082) 正常运行

## 项目结构

```
python/
├── app.py                      # Flask 应用入口
├── config.py                   # 配置文件
├── models.py                   # 数据库模型
├── routes_borrows.py           # 借用记录路由
├── requirements.txt            # 依赖列表
├── services/
│   ├── user_client.py          # 用户服务客户端
│   └── material_client.py      # 物资服务客户端
└── utils/
    └── response.py             # 统一响应格式工具
```

## 快速开始

### 1. 安装依赖

```bash
cd python
pip install -r requirements.txt
```

### 2. 配置数据库

在 MySQL 中创建数据库:

```sql
CREATE DATABASE borrow_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 3. 配置环境变量（可选）

创建 `.env` 文件（或使用默认配置）:

```env
# 服务器配置
HOST=0.0.0.0
PORT=8081
DEBUG=True

# 数据库配置
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=borrow_db

# 其他服务地址
USER_SERVICE_BASE_URL=http://localhost:8083
MATERIAL_SERVICE_BASE_URL=http://localhost:8082
```

### 4. 启动服务

```bash
python app.py
```

正常启动输出如下:

```
借用记录服务启动中...
监听地址: 0.0.0.0:8081
数据库: mysql+pymysql://root:***@localhost:3306/borrow_db?charset=utf8mb4
用户服务: http://localhost:8083
物资服务: http://localhost:8082
数据库表创建成功!
 * Debug mode: on
WARNING: This is a development server. Do not use it in a production deployment.
 * Running on http://0.0.0.0:8081
Press CTRL+C to quit
```

## API 接口说明

### Swagger 在线文档

- 启动服务后访问 `http://localhost:8081/apidocs` 查看自动生成的 Swagger UI
- 文档基于 [Flasgger](https://github.com/flasgger/flasgger)，会展示接口请求参数、响应结构、示例等
- 在 Swagger UI 中可以直接发起请求（需确保后台服务及依赖服务均可访问）

### 健康检查

- `GET /` - 服务信息
- `GET /health` - 健康检查

### 借用记录管理

| 功能           | 方法   | 路径                    | 说明                     |
|----------------|--------|-------------------------|--------------------------|
| 创建借用记录   | POST   | `/borrows`              | 借出操作（创建记录）     |
| 查询借用列表   | GET    | `/borrows`              | 支持按用户/物资/状态过滤 |
| 查询单条记录   | GET    | `/borrows/{id}`         | 根据借用记录 ID 查询     |
| 更新借用记录   | PUT    | `/borrows/{id}`         | 修改备注、归还时间、状态 |
| 删除借用记录   | DELETE | `/borrows/{id}`         | 删除记录                 |
| 归还操作       | POST   | `/borrows/{id}/return`  | 专门的归还接口           |

### 示例请求

#### 1. 创建借用记录（借出）

```bash
POST http://localhost:8081/borrows
Content-Type: application/json

{
  "userId": 1,
  "materialId": 1001,
  "quantity": 1,
  "dueAt": "2025-11-27T10:30:00",
  "remark": "实验1使用"
}
```

#### 2. 查询借用列表

```bash
# 查询所有
GET http://localhost:8081/borrows

# 查询某用户的借用记录
GET http://localhost:8081/borrows?userId=1

# 查询某物资的借用记录
GET http://localhost:8081/borrows?materialId=1001

# 查询借出中的记录
GET http://localhost:8081/borrows?status=0

# 包含用户和物资详细信息
GET http://localhost:8081/borrows?include=user,material

# 分页查询
GET http://localhost:8081/borrows?page=1&pageSize=10
```

#### 3. 查询单条记录

```bash
GET http://localhost:8081/borrows/1
```

#### 4. 归还操作

```bash
POST http://localhost:8081/borrows/1/return
Content-Type: application/json

{
  "returnedAt": "2025-11-21T10:30:00",
  "remark": "正常归还"
}
```

#### 5. 更新借用记录

```bash
PUT http://localhost:8081/borrows/1
Content-Type: application/json

{
  "status": 1,
  "returnedAt": "2025-11-21T10:30:00",
  "remark": "提前归还"
}
```

#### 6. 删除借用记录

```bash
DELETE http://localhost:8081/borrows/1
```

## 数据库表结构

### borrow_records 表

| 字段名        | 类型         | 说明                              |
|--------------|--------------|-----------------------------------|
| id           | BIGINT       | 主键，借用记录 ID                 |
| user_id      | BIGINT       | 借用人 ID                         |
| material_id  | BIGINT       | 物资 ID                           |
| quantity     | INT          | 借用数量                          |
| status       | SMALLINT     | 状态: 0-借出中, 1-已归还, 2-已取消 |
| borrowed_at  | DATETIME     | 借出时间                          |
| due_at       | DATETIME     | 应归还时间                        |
| returned_at  | DATETIME     | 实际归还时间                      |
| remark       | VARCHAR(255) | 备注信息                          |
| created_at   | DATETIME     | 创建时间                          |
| updated_at   | DATETIME     | 更新时间                          |

## 统一响应格式

所有接口返回统一格式:

```json
{
  "code": 200,
  "message": "success",
  "data": { },
  "timestamp": "2025-11-20T15:00:00"
}
```

**状态码说明:**
- `200`: 成功
- `201`: 创建成功
- `400`: 参数错误
- `404`: 资源不存在
- `409`: 业务冲突
- `500`: 服务器内部错误

## 开发说明

### 目录说明

- `app.py`: Flask 应用入口，初始化数据库和注册路由
- `config.py`: 配置管理，包含数据库和其他服务地址
- `models.py`: SQLAlchemy 数据库模型定义
- `routes_borrows.py`: 借用记录相关的路由和业务逻辑
- `services/`: 服务间调用的客户端封装
- `utils/`: 工具函数，如统一响应格式

### 与其他服务的交互

1. **人员管理服务 (8083)**: 验证用户是否存在
2. **物资管理服务 (8082)**: 检查物资可用性，更新物资状态

## 注意事项

1. 首次运行时会自动创建数据库表
2. 确保 MySQL 数据库已创建并正确配置
3. 确保人员管理服务和物资管理服务已启动
4. 生产环境请使用 WSGI 服务器（如 Gunicorn）而非 Flask 自带的开发服务器

## 故障排查

### 数据库连接失败

检查 MySQL 是否运行，数据库配置是否正确:
```bash
mysql -u root -p -e "SHOW DATABASES;"
```

### 服务间调用失败

检查其他服务是否正常运行:
```bash
curl http://localhost:8083/users/1
curl http://localhost:8082/materials/1
```

## 参考文档

详细的接口设计规范请参考项目根目录下的 `borrow-service-spec.md` 文件。
