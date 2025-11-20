package com.example.demo.entity;

import java.time.LocalDateTime;

import io.swagger.v3.oas.annotations.media.Schema;

/**
 * 用户实体类
 * 对应数据库表：users
 */
@Schema(name = "User", description = "用户基础信息及其关联的角色数据")
public class User {
    /**
     * 用户ID
     */
    @Schema(description = "主键ID", example = "1")
    private Long id;

    /**
     * 姓名
     */
    @Schema(description = "用户姓名", example = "张三")
    private String name;

    /**
     * 角色ID（外键）
     */
    @Schema(description = "关联的角色ID", example = "1001")
    private Long roleId;

    /**
     * 创建时间
     */
    @Schema(description = "创建时间")
    private LocalDateTime createdAt;

    /**
     * 更新时间
     */
    @Schema(description = "最近更新时间")
    private LocalDateTime updatedAt;

    /**
     * 关联的角色对象（用于查询时关联）
     */
    @Schema(description = "角色详情")
    private Role role;

    // 无参构造函数
    public User() {
    }

    // 有参构造函数
    public User(String name, Long roleId) {
        this.name = name;
        this.roleId = roleId;
    }

    // Getter和Setter方法
    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Long getRoleId() {
        return roleId;
    }

    public void setRoleId(Long roleId) {
        this.roleId = roleId;
    }

    public LocalDateTime getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(LocalDateTime createdAt) {
        this.createdAt = createdAt;
    }

    public LocalDateTime getUpdatedAt() {
        return updatedAt;
    }

    public void setUpdatedAt(LocalDateTime updatedAt) {
        this.updatedAt = updatedAt;
    }

    public Role getRole() {
        return role;
    }

    public void setRole(Role role) {
        this.role = role;
    }

    @Override
    public String toString() {
        return "User{" +
                "id=" + id +
                ", name='" + name + '\'' +
                ", roleId=" + roleId +
                ", createdAt=" + createdAt +
                ", updatedAt=" + updatedAt +
                ", role=" + role +
                '}';
    }
}

