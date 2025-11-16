package com.example.demo.mapper;

import com.example.demo.entity.Role;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;

import java.util.List;

/**
 * 角色Mapper接口
 * 对应数据库表：roles
 */
@Mapper
public interface RoleMapper {

    /**
     * 查询所有角色
     * @return 角色列表
     */
    List<Role> findAll();

    /**
     * 根据ID查询角色
     * @param id 角色ID
     * @return 角色对象
     */
    Role findById(@Param("id") Long id);

    /**
     * 根据角色名称查询角色
     * @param roleName 角色名称
     * @return 角色对象
     */
    Role findByRoleName(@Param("roleName") String roleName);

    /**
     * 插入角色
     * @param role 角色对象
     * @return 影响行数
     */
    int insert(Role role);

    /**
     * 更新角色
     * @param role 角色对象
     * @return 影响行数
     */
    int update(Role role);

    /**
     * 根据ID删除角色
     * @param id 角色ID
     * @return 影响行数
     */
    int deleteById(@Param("id") Long id);
}

