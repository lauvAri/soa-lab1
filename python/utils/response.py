from flask import jsonify
from datetime import datetime


def make_response(code=200, message="success", data=None):
    """
    统一响应格式封装
    
    Args:
        code: 业务状态码
        message: 人类可读信息
        data: 具体数据对象或数组
    
    Returns:
        (response, http_status_code)
    """
    return jsonify({
        "code": code,
        "message": message,
        "data": data,
        "timestamp": datetime.utcnow().isoformat()
    }), code if code < 600 else 500


def success_response(data=None, message="success", code=200):
    """成功响应"""
    return make_response(code=code, message=message, data=data)


def created_response(data=None, message="创建成功"):
    """创建成功响应 (201)"""
    return make_response(code=201, message=message, data=data)


def bad_request_response(message="请求参数不合法"):
    """参数错误响应 (400)"""
    return make_response(code=400, message=message, data=None)


def not_found_response(message="资源不存在"):
    """资源不存在响应 (404)"""
    return make_response(code=404, message=message, data=None)


def conflict_response(message="业务冲突"):
    """业务冲突响应 (409)"""
    return make_response(code=409, message=message, data=None)


def internal_error_response(message="服务器内部错误"):
    """服务器内部错误响应 (500)"""
    return make_response(code=500, message=message, data=None)
