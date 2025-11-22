from flask import Blueprint, request
from models import db, BorrowRecord
from utils.response import (
    success_response, created_response, bad_request_response,
    not_found_response, conflict_response, internal_error_response
)
from services.user_client import UserClient
from services.material_client import MaterialClient
from datetime import datetime
from config import Config

# 创建蓝图
borrows_bp = Blueprint('borrows', __name__)

# 初始化客户端
user_client = UserClient()
material_client = MaterialClient()


@borrows_bp.route('/borrows', methods=['POST'])
def create_borrow():
    """
    创建借用记录（借出）
    POST /borrows
    """
    try:
        # 获取请求数据
        data = request.get_json()
        if not data:
            return bad_request_response("请求体不能为空")
        
        # 校验必填参数
        user_id = data.get('userId')
        material_id = data.get('materialId')
        
        if not user_id or not isinstance(user_id, int) or user_id <= 0:
            return bad_request_response("userId 必填且必须为正整数")
        
        if not material_id or not isinstance(material_id, int) or material_id <= 0:
            return bad_request_response("materialId 必填且必须为正整数")
        
        # 获取可选参数
        quantity = data.get('quantity', 1)
        due_at_str = data.get('dueAt')
        remark = data.get('remark', '')
        
        # 校验数量
        if not isinstance(quantity, int) or quantity <= 0:
            return bad_request_response("quantity 必须为正整数")
        
        # 解析应归还时间
        due_at = None
        if due_at_str:
            try:
                due_at = datetime.fromisoformat(due_at_str.replace('Z', '+00:00'))
            except ValueError:
                return bad_request_response("dueAt 时间格式不正确，应为 ISO 8601 格式")
        
        # 1. 验证用户存在
        try:
            if not user_client.check_user_exists(user_id):
                return not_found_response("用户不存在")
        except Exception as e:
            return internal_error_response(f"调用用户服务失败: {str(e)}")
        
        # 2. 验证物资存在且可用
        try:
            is_available, material_data = material_client.check_material_available(material_id)
            if material_data is None:
                return not_found_response("物资不存在")
            if not is_available:
                return conflict_response("物资当前不可借出")
        except Exception as e:
            return internal_error_response(f"调用物资服务失败: {str(e)}")
        
        # 3. 创建借用记录
        borrow_record = BorrowRecord(
            user_id=user_id,
            material_id=material_id,
            quantity=quantity,
            status=BorrowRecord.STATUS_BORROWED,
            borrowed_at=datetime.utcnow(),
            due_at=due_at,
            remark=remark
        )
        
        db.session.add(borrow_record)
        db.session.commit()
        
        # 4. 更新物资状态为借出中
        try:
            material_client.mark_as_borrowed(material_id)
        except Exception as e:
            # 如果更新物资状态失败，回滚借用记录
            db.session.delete(borrow_record)
            db.session.commit()
            return internal_error_response(f"更新物资状态失败: {str(e)}")
        
        # 返回创建的记录
        return created_response(
            data=borrow_record.to_dict(),
            message="创建借用记录成功"
        )
    
    except Exception as e:
        db.session.rollback()
        return internal_error_response(f"创建借用记录失败: {str(e)}")


@borrows_bp.route('/borrows', methods=['GET'])
def list_borrows():
    """
    查询借用列表（支持过滤和分页）
    GET /borrows?userId=1&materialId=1001&status=0&page=1&pageSize=10&include=user,material
    """
    try:
        # 获取查询参数
        user_id = request.args.get('userId', type=int)
        material_id = request.args.get('materialId', type=int)
        status = request.args.get('status', type=int)
        page = request.args.get('page', 1, type=int)
        page_size = request.args.get('pageSize', Config.DEFAULT_PAGE_SIZE, type=int)
        include = request.args.get('include', '')
        
        # 限制分页大小
        if page < 1:
            page = 1
        if page_size < 1 or page_size > Config.MAX_PAGE_SIZE:
            page_size = Config.DEFAULT_PAGE_SIZE
        
        # 构建查询
        query = BorrowRecord.query
        
        if user_id is not None:
            query = query.filter_by(user_id=user_id)
        
        if material_id is not None:
            query = query.filter_by(material_id=material_id)
        
        if status is not None:
            if status not in [0, 1, 2]:
                return bad_request_response("status 必须为 0, 1 或 2")
            query = query.filter_by(status=status)
        
        # 按创建时间倒序排列
        query = query.order_by(BorrowRecord.created_at.desc())
        
        # 分页
        pagination = query.paginate(page=page, per_page=page_size, error_out=False)
        
        # 是否包含用户和物资信息
        include_user = 'user' in include.lower()
        include_material = 'material' in include.lower()
        
        # 转换为字典列表
        items = []
        for record in pagination.items:
            user_data = None
            material_data = None
            
            # 如果需要包含用户信息
            if include_user:
                try:
                    user_data = user_client.get_user(record.user_id)
                except:
                    pass  # 忽略获取用户信息失败
            
            # 如果需要包含物资信息
            if include_material:
                try:
                    material_data = material_client.get_material(record.material_id)
                except:
                    pass  # 忽略获取物资信息失败
            
            items.append(record.to_dict(
                include_user=include_user,
                include_material=include_material,
                user_data=user_data,
                material_data=material_data
            ))
        
        # 返回结果
        return success_response(data={
            'items': items,
            'page': page,
            'pageSize': page_size,
            'total': pagination.total
        })
    
    except Exception as e:
        return internal_error_response(f"查询借用列表失败: {str(e)}")


@borrows_bp.route('/borrows/<int:id>', methods=['GET'])
def get_borrow(id):
    """
    查询单条借用记录
    GET /borrows/{id}
    """
    try:
        # 查询记录
        record = BorrowRecord.query.get(id)
        if not record:
            return not_found_response("借用记录不存在")
        
        # 获取查询参数
        include = request.args.get('include', '')
        include_user = 'user' in include.lower()
        include_material = 'material' in include.lower()
        
        user_data = None
        material_data = None
        
        # 如果需要包含用户信息
        if include_user:
            try:
                user_data = user_client.get_user(record.user_id)
            except:
                pass
        
        # 如果需要包含物资信息
        if include_material:
            try:
                material_data = material_client.get_material(record.material_id)
            except:
                pass
        
        return success_response(data=record.to_dict(
            include_user=include_user,
            include_material=include_material,
            user_data=user_data,
            material_data=material_data
        ))
    
    except Exception as e:
        return internal_error_response(f"查询借用记录失败: {str(e)}")


@borrows_bp.route('/borrows/<int:id>', methods=['PUT'])
def update_borrow(id):
    """
    更新借用记录（包括归还）
    PUT /borrows/{id}
    """
    try:
        # 查询记录
        record = BorrowRecord.query.get(id)
        if not record:
            return not_found_response("借用记录不存在")
        
        # 获取请求数据
        data = request.get_json()
        if not data:
            return bad_request_response("请求体不能为空")
        
        # 获取要更新的字段
        new_status = data.get('status')
        returned_at_str = data.get('returnedAt')
        due_at_str = data.get('dueAt')
        remark = data.get('remark')
        
        # 处理状态变更（归还操作）
        if new_status is not None:
            if new_status not in [0, 1, 2]:
                return bad_request_response("status 必须为 0, 1 或 2")
            
            # 从借出中到已归还
            if record.status == BorrowRecord.STATUS_BORROWED and new_status == BorrowRecord.STATUS_RETURNED:
                # 设置归还时间
                if returned_at_str:
                    try:
                        record.returned_at = datetime.fromisoformat(returned_at_str.replace('Z', '+00:00'))
                    except ValueError:
                        return bad_request_response("returnedAt 时间格式不正确")
                else:
                    record.returned_at = datetime.utcnow()
                
                # 更新物资状态为可用
                try:
                    material_client.mark_as_available(record.material_id)
                except Exception as e:
                    return internal_error_response(f"更新物资状态失败: {str(e)}")
            
            record.status = new_status
        
        # 更新应归还时间
        if due_at_str:
            try:
                record.due_at = datetime.fromisoformat(due_at_str.replace('Z', '+00:00'))
            except ValueError:
                return bad_request_response("dueAt 时间格式不正确")
        
        # 更新备注
        if remark is not None:
            record.remark = remark
        
        # 更新时间戳
        record.updated_at = datetime.utcnow()
        
        db.session.commit()
        
        return success_response(
            data=record.to_dict(),
            message="借用记录更新成功"
        )
    
    except Exception as e:
        db.session.rollback()
        return internal_error_response(f"更新借用记录失败: {str(e)}")


@borrows_bp.route('/borrows/<int:id>/return', methods=['POST'])
def return_borrow(id):
    """
    归还操作专用接口
    POST /borrows/{id}/return
    """
    try:
        # 查询记录
        record = BorrowRecord.query.get(id)
        if not record:
            return not_found_response("借用记录不存在")
        
        # 检查当前状态
        if record.status != BorrowRecord.STATUS_BORROWED:
            return conflict_response("借用记录当前状态不允许归还操作")
        
        # 获取请求数据
        data = request.get_json() or {}
        returned_at_str = data.get('returnedAt')
        remark = data.get('remark')
        
        # 设置归还时间
        if returned_at_str:
            try:
                record.returned_at = datetime.fromisoformat(returned_at_str.replace('Z', '+00:00'))
            except ValueError:
                return bad_request_response("returnedAt 时间格式不正确")
        else:
            record.returned_at = datetime.utcnow()
        
        # 更新备注
        if remark is not None:
            record.remark = remark
        
        # 更新状态为已归还
        record.status = BorrowRecord.STATUS_RETURNED
        record.updated_at = datetime.utcnow()
        
        # 更新物资状态为可用
        try:
            material_client.mark_as_available(record.material_id)
        except Exception as e:
            return internal_error_response(f"更新物资状态失败: {str(e)}")
        
        db.session.commit()
        
        return success_response(
            data=record.to_dict(),
            message="归还成功"
        )
    
    except Exception as e:
        db.session.rollback()
        return internal_error_response(f"归还操作失败: {str(e)}")


@borrows_bp.route('/borrows/<int:id>', methods=['DELETE'])
def delete_borrow(id):
    """
    删除借用记录
    DELETE /borrows/{id}
    """
    try:
        # 查询记录
        record = BorrowRecord.query.get(id)
        if not record:
            return not_found_response("借用记录不存在")
        
        # 如果记录状态为借出中，先归还
        if record.status == BorrowRecord.STATUS_BORROWED:
            try:
                material_client.mark_as_available(record.material_id)
            except Exception as e:
                return internal_error_response(f"更新物资状态失败: {str(e)}")
        
        # 删除记录
        db.session.delete(record)
        db.session.commit()
        
        return success_response(
            data=None,
            message="借用记录删除成功"
        )
    
    except Exception as e:
        db.session.rollback()
        return internal_error_response(f"删除借用记录失败: {str(e)}")
