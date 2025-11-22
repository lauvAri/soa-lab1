import requests
from config import Config


class UserClient:
    """人员管理服务客户端 (Java / 8083)"""
    
    def __init__(self):
        self.base_url = Config.USER_SERVICE_BASE_URL
    
    def get_user(self, user_id):
        """
        获取用户信息
        
        Args:
            user_id: 用户ID
        
        Returns:
            dict: 用户信息，如果成功返回 {'id', 'name', 'roleId'}
            None: 如果用户不存在或服务不可用
        
        Raises:
            Exception: 服务调用失败
        """
        try:
            url = f"{self.base_url}/users/{user_id}"
            response = requests.get(url, timeout=5)
            
            if response.status_code == 200:
                result = response.json()
                # Java 服务返回格式: {"code": 200, "message": "success", "data": {...}}
                if result.get('code') == 200 and result.get('data'):
                    return result['data']
                return None
            elif response.status_code == 404:
                return None
            else:
                raise Exception(f"用户服务返回错误: {response.status_code}")
        
        except requests.exceptions.Timeout:
            raise Exception("用户服务调用超时")
        except requests.exceptions.ConnectionError:
            raise Exception("无法连接到用户服务")
        except Exception as e:
            raise Exception(f"调用用户服务失败: {str(e)}")
    
    def check_user_exists(self, user_id):
        """
        检查用户是否存在
        
        Args:
            user_id: 用户ID
        
        Returns:
            bool: 用户是否存在
        """
        try:
            user = self.get_user(user_id)
            return user is not None
        except:
            return False
