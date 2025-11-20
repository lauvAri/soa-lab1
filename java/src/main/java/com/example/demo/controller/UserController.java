package com.example.demo.controller;

import com.example.demo.common.ResponseResult;
import com.example.demo.entity.User;
import com.example.demo.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.Parameter;
import io.swagger.v3.oas.annotations.tags.Tag;

/**
 * 用户控制器
 * 提供用户管理的REST API接口
 */
@Tag(name = "用户管理", description = "提供用户查询、角色筛选、增删改等接口")
@RestController
@RequestMapping("/users")
public class UserController {

    @Autowired
    private UserService userService;

    /**
     * 获取所有用户
     * GET /users
     * @return 用户列表（包含角色信息）
     */
    @Operation(summary = "获取全部用户", description = "返回所有用户以及关联的角色信息")
    @GetMapping
    public ResponseEntity<ResponseResult<List<User>>> getAllUsers() {
        List<User> users = userService.getAllUsers();
        return ResponseEntity.ok(ResponseResult.success(users));
    }

    /**
     * 根据角色ID获取用户列表
     * GET /users/role/{roleId}
     * 注意：这个路径必须放在 /users/{id} 之前，避免路径冲突
     * @param roleId 角色ID
     * @return 用户列表
     */
    @Operation(summary = "按角色查询用户", description = "根据角色ID获取所属用户列表")
    @GetMapping("/role/{roleId}")
    public ResponseEntity<ResponseResult<List<User>>> getUsersByRoleId(
            @Parameter(description = "角色ID", example = "1001")
            @PathVariable Long roleId) {
        List<User> users = userService.getUsersByRoleId(roleId);
        return ResponseEntity.ok(ResponseResult.success(users));
    }

    /**
     * 根据ID获取用户
     * GET /users/{id}
     * @param id 用户ID
     * @return 用户对象（包含角色信息）
     */
    @Operation(summary = "根据ID获取用户", description = "根据用户ID返回其详情及角色")
    @GetMapping("/{id}")
    public ResponseEntity<ResponseResult<User>> getUserById(
            @Parameter(description = "用户ID", example = "1")
            @PathVariable Long id) {
        User user = userService.getUserById(id);
        return ResponseEntity.ok(ResponseResult.success(user));
    }

    /**
     * 创建用户
     * POST /users
     * @param user 用户对象（需要name和roleId）
     * @return 创建的用户对象（包含角色信息）
     */
    @Operation(summary = "创建用户", description = "根据姓名与角色ID创建新用户")
    @PostMapping
    public ResponseEntity<ResponseResult<User>> createUser(
            @io.swagger.v3.oas.annotations.parameters.RequestBody(
                    description = "待创建的用户信息", required = true)
            @RequestBody User user) {
        // 验证必填字段
        if (user.getName() == null || user.getName().trim().isEmpty()) {
            return ResponseEntity.badRequest()
                    .body(ResponseResult.badRequest("姓名不能为空"));
        }
        if (user.getRoleId() == null) {
            return ResponseEntity.badRequest()
                    .body(ResponseResult.badRequest("角色ID不能为空"));
        }

        User createdUser = userService.createUser(user);
        return ResponseEntity.status(HttpStatus.CREATED)
                .body(ResponseResult.success("用户创建成功", createdUser));
    }

    /**
     * 更新用户
     * PUT /users/{id}
     * @param id 用户ID
     * @param user 用户对象（需要name和roleId）
     * @return 更新后的用户对象（包含角色信息）
     */
    @Operation(summary = "更新用户", description = "根据ID修改用户姓名与角色信息")
    @PutMapping("/{id}")
    public ResponseEntity<ResponseResult<User>> updateUser(
            @Parameter(description = "用户ID", example = "1")
            @PathVariable Long id,
            @io.swagger.v3.oas.annotations.parameters.RequestBody(
                    description = "新的用户信息", required = true)
            @RequestBody User user) {
        // 验证必填字段
        if (user.getName() == null || user.getName().trim().isEmpty()) {
            return ResponseEntity.badRequest()
                    .body(ResponseResult.badRequest("姓名不能为空"));
        }
        if (user.getRoleId() == null) {
            return ResponseEntity.badRequest()
                    .body(ResponseResult.badRequest("角色ID不能为空"));
        }

        User updatedUser = userService.updateUser(id, user);
        return ResponseEntity.ok(ResponseResult.success("用户更新成功", updatedUser));
    }

    /**
     * 删除用户
     * DELETE /users/{id}
     * @param id 用户ID
     * @return 删除结果
     */
    @Operation(summary = "删除用户", description = "根据用户ID删除用户")
    @DeleteMapping("/{id}")
    public ResponseEntity<ResponseResult<Object>> deleteUser(
            @Parameter(description = "用户ID", example = "1")
            @PathVariable Long id) {
        userService.deleteUser(id);
        return ResponseEntity.ok(ResponseResult.success("用户删除成功"));
    }
}
