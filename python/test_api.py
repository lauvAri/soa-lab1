"""
借用记录服务测试脚本
用于测试基本 API 功能
"""
import requests
import json
from datetime import datetime, timedelta

BASE_URL = "http://localhost:8081"


def test_health_check():
    """测试健康检查接口"""
    print("\n=== 测试健康检查 ===")
    response = requests.get(f"{BASE_URL}/")
    print(f"状态码: {response.status_code}")
    print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
    return response.status_code == 200


def test_create_borrow():
    """测试创建借用记录"""
    print("\n=== 测试创建借用记录 ===")
    
    due_date = (datetime.now() + timedelta(days=7)).isoformat()
    
    data = {
        "userId": 1,
        "materialId": 1,
        "quantity": 1,
        "dueAt": due_date,
        "remark": "测试借用"
    }
    
    response = requests.post(
        f"{BASE_URL}/borrows",
        json=data,
        headers={"Content-Type": "application/json"}
    )
    
    print(f"状态码: {response.status_code}")
    print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
    
    if response.status_code == 201:
        return response.json()['data']['id']
    return None


def test_list_borrows():
    """测试查询借用列表"""
    print("\n=== 测试查询借用列表 ===")
    
    response = requests.get(f"{BASE_URL}/borrows")
    print(f"状态码: {response.status_code}")
    result = response.json()
    print(f"总记录数: {result['data']['total']}")
    print(f"响应: {json.dumps(result, indent=2, ensure_ascii=False)}")
    return response.status_code == 200


def test_get_borrow(borrow_id):
    """测试查询单条借用记录"""
    print(f"\n=== 测试查询借用记录 {borrow_id} ===")
    
    response = requests.get(f"{BASE_URL}/borrows/{borrow_id}")
    print(f"状态码: {response.status_code}")
    print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
    return response.status_code == 200


def test_return_borrow(borrow_id):
    """测试归还操作"""
    print(f"\n=== 测试归还借用记录 {borrow_id} ===")
    
    data = {
        "returnedAt": datetime.now().isoformat(),
        "remark": "测试归还"
    }
    
    response = requests.post(
        f"{BASE_URL}/borrows/{borrow_id}/return",
        json=data,
        headers={"Content-Type": "application/json"}
    )
    
    print(f"状态码: {response.status_code}")
    print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
    return response.status_code == 200


def test_delete_borrow(borrow_id):
    """测试删除借用记录"""
    print(f"\n=== 测试删除借用记录 {borrow_id} ===")
    
    response = requests.delete(f"{BASE_URL}/borrows/{borrow_id}")
    print(f"状态码: {response.status_code}")
    print(f"响应: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
    return response.status_code == 200


def main():
    """运行所有测试"""
    print("=" * 50)
    print("借用记录服务 API 测试")
    print("=" * 50)
    
    try:
        # 1. 健康检查
        if not test_health_check():
            print("\n❌ 健康检查失败，请确保服务已启动")
            return
        
        print("\n✓ 健康检查通过")
        
        # 2. 查询列表
        test_list_borrows()
        
        # 3. 创建借用记录（需要用户和物资服务支持）
        print("\n注意: 以下测试需要用户服务(8083)和物资服务(8082)正常运行")
        print("如果失败，请确保这些服务已启动且有对应的测试数据")
        
        # borrow_id = test_create_borrow()
        
        # if borrow_id:
        #     # 4. 查询单条记录
        #     test_get_borrow(borrow_id)
        #     
        #     # 5. 归还操作
        #     test_return_borrow(borrow_id)
        #     
        #     # 6. 删除记录
        #     test_delete_borrow(borrow_id)
        
        print("\n" + "=" * 50)
        print("✓ 基本测试完成")
        print("=" * 50)
    
    except requests.exceptions.ConnectionError:
        print("\n❌ 无法连接到服务，请确保服务已在 8081 端口启动")
    except Exception as e:
        print(f"\n❌ 测试失败: {str(e)}")


if __name__ == '__main__':
    main()
