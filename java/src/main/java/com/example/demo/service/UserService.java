package com.example.demo.service;

import com.example.demo.entity.User;
import com.example.demo.exception.RoleNotFoundException;
import com.example.demo.exception.UserNotFoundException;
import com.example.demo.mapper.RoleMapper;
import com.example.demo.mapper.UserMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

/**
 * 用户服务类
 * 提供用户相关的业务逻辑
 */
@Service
public class UserService {

    @Autowired
    private UserMapper userMapper;

    @Autowired
    private RoleMapper roleMapper;

    /**
     * 获取所有用户（包含角色信息）
     * @return 用户列表
     */
    public List<User> getAllUsers() {
        return userMapper.findAllWithRole();
    }

    /**
     * 根据ID获取用户（包含角色信息）
     * @param id 用户ID
     * @return 用户对象
     * @throws UserNotFoundException 用户不存在时抛出
     */
    public User getUserById(Long id) {
        User user = userMapper.findByIdWithRole(id);
        if (user == null) {
            throw new UserNotFoundException(id);
        }
        return user;
    }

    /**
     * 创建用户
     * @param user 用户对象
     * @return 创建的用户对象（包含角色信息）
     * @throws RoleNotFoundException 角色不存在时抛出
     */
    @Transactional
    public User createUser(User user) {
        // 验证角色是否存在
        if (user.getRoleId() == null) {
            throw new IllegalArgumentException("角色ID不能为空");
        }
        if (roleMapper.findById(user.getRoleId()) == null) {
            throw new RoleNotFoundException(user.getRoleId());
        }

        // 插入用户
        userMapper.insert(user);
        
        // 返回包含角色信息的用户对象
        return userMapper.findByIdWithRole(user.getId());
    }

    /**
     * 更新用户
     * @param id 用户ID
     * @param user 用户对象
     * @return 更新后的用户对象（包含角色信息）
     * @throws UserNotFoundException 用户不存在时抛出
     * @throws RoleNotFoundException 角色不存在时抛出
     */
    @Transactional
    public User updateUser(Long id, User user) {
        // 验证用户是否存在
        if (userMapper.findById(id) == null) {
            throw new UserNotFoundException(id);
        }

        // 验证角色是否存在
        if (user.getRoleId() != null && roleMapper.findById(user.getRoleId()) == null) {
            throw new RoleNotFoundException(user.getRoleId());
        }

        // 设置用户ID
        user.setId(id);
        
        // 更新用户
        userMapper.update(user);
        
        // 返回更新后的用户对象（包含角色信息）
        return userMapper.findByIdWithRole(id);
    }

    /**
     * 删除用户
     * @param id 用户ID
     * @throws UserNotFoundException 用户不存在时抛出
     */
    @Transactional
    public void deleteUser(Long id) {
        // 验证用户是否存在
        if (userMapper.findById(id) == null) {
            throw new UserNotFoundException(id);
        }
        
        // 删除用户
        userMapper.deleteById(id);
    }

    /**
     * 根据角色ID获取用户列表
     * @param roleId 角色ID
     * @return 用户列表
     */
    public List<User> getUsersByRoleId(Long roleId) {
        return userMapper.findByRoleId(roleId);
    }
}

