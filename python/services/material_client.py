import requests
from config import Config


class MaterialClient:
    """物资管理服务客户端 (Go / 8082)"""
    
    def __init__(self):
        self.base_url = Config.MATERIAL_SERVICE_BASE_URL
    
    def get_material(self, material_id):
        """
        获取物资信息
        
        Args:
            material_id: 物资ID
        
        Returns:
            dict: 物资信息，如果成功返回 {'materialId', 'materialName', 'materialStatus', ...}
            None: 如果物资不存在或服务不可用
        
        Raises:
            Exception: 服务调用失败
        """
        try:
            url = f"{self.base_url}/materials/{material_id}"
            response = requests.get(url, timeout=5)
            
            if response.status_code == 200:
                result = response.json()
                # Go 服务返回格式: {"code": 200, "message": "success", "data": {...}}
                if result.get('code') == 200 and result.get('data'):
                    return result['data']
                return None
            elif response.status_code == 404:
                return None
            else:
                raise Exception(f"物资服务返回错误: {response.status_code}")
        
        except requests.exceptions.Timeout:
            raise Exception("物资服务调用超时")
        except requests.exceptions.ConnectionError:
            raise Exception("无法连接到物资服务")
        except Exception as e:
            raise Exception(f"调用物资服务失败: {str(e)}")
    
    def check_material_available(self, material_id):
        """
        检查物资是否可借（存在且状态为可用 status == 0）
        
        Args:
            material_id: 物资ID
        
        Returns:
            tuple: (is_available, material_data)
                is_available: bool, 是否可借
                material_data: dict, 物资信息（如果存在）
        """
        try:
            material = self.get_material(material_id)
            if material is None:
                return False, None
            
            # material_status: 0-可用, 1-借出中, 2-维护中
            is_available = material.get('materialStatus') == 0
            return is_available, material
        except:
            return False, None
    
    def update_material_status(self, material_id, status):
        """
        更新物资状态
        
        Args:
            material_id: 物资ID
            status: 物资状态 (0-可用, 1-借出中, 2-维护中)
        
        Returns:
            bool: 是否更新成功
        
        Raises:
            Exception: 服务调用失败
        """
        try:
            url = f"{self.base_url}/materials/{material_id}"
            payload = {"materialStatus": status}
            response = requests.put(url, json=payload, timeout=5)
            
            if response.status_code == 200:
                result = response.json()
                return result.get('code') == 200
            else:
                raise Exception(f"物资服务返回错误: {response.status_code}")
        
        except requests.exceptions.Timeout:
            raise Exception("物资服务调用超时")
        except requests.exceptions.ConnectionError:
            raise Exception("无法连接到物资服务")
        except Exception as e:
            raise Exception(f"更新物资状态失败: {str(e)}")
    
    def mark_as_borrowed(self, material_id):
        """
        标记物资为借出中 (status = 1)
        
        Args:
            material_id: 物资ID
        
        Returns:
            bool: 是否更新成功
        """
        return self.update_material_status(material_id, 1)
    
    def mark_as_available(self, material_id):
        """
        标记物资为可用 (status = 0)
        
        Args:
            material_id: 物资ID
        
        Returns:
            bool: 是否更新成功
        """
        return self.update_material_status(material_id, 0)
