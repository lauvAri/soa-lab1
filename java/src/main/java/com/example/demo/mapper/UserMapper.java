package com.example.demo.mapper;

import com.example.demo.entity.User;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;

import java.util.List;

/**
 * 用户Mapper接口
 * 对应数据库表：users
 */
@Mapper
public interface UserMapper {

    /**
     * 查询所有用户（不包含角色信息）
     * @return 用户列表
     */
    List<User> findAll();

    /**
     * 查询所有用户（包含角色信息）
     * @return 用户列表（包含角色对象）
     */
    List<User> findAllWithRole();

    /**
     * 根据ID查询用户（不包含角色信息）
     * @param id 用户ID
     * @return 用户对象
     */
    User findById(@Param("id") Long id);

    /**
     * 根据ID查询用户（包含角色信息）
     * @param id 用户ID
     * @return 用户对象（包含角色对象）
     */
    User findByIdWithRole(@Param("id") Long id);

    /**
     * 根据角色ID查询用户列表
     * @param roleId 角色ID
     * @return 用户列表
     */
    List<User> findByRoleId(@Param("roleId") Long roleId);

    /**
     * 插入用户
     * @param user 用户对象
     * @return 影响行数
     */
    int insert(User user);

    /**
     * 更新用户
     * @param user 用户对象
     * @return 影响行数
     */
    int update(User user);

    /**
     * 根据ID删除用户
     * @param id 用户ID
     * @return 影响行数
     */
    int deleteById(@Param("id") Long id);
}

