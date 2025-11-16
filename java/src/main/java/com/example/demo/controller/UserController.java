package com.example.demo.controller;

import com.example.demo.common.ResponseResult;
import com.example.demo.entity.User;
import com.example.demo.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

/**
 * 用户控制器
 * 提供用户管理的REST API接口
 */
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
    @GetMapping("/role/{roleId}")
    public ResponseEntity<ResponseResult<List<User>>> getUsersByRoleId(@PathVariable Long roleId) {
        List<User> users = userService.getUsersByRoleId(roleId);
        return ResponseEntity.ok(ResponseResult.success(users));
    }

    /**
     * 根据ID获取用户
     * GET /users/{id}
     * @param id 用户ID
     * @return 用户对象（包含角色信息）
     */
    @GetMapping("/{id}")
    public ResponseEntity<ResponseResult<User>> getUserById(@PathVariable Long id) {
        User user = userService.getUserById(id);
        return ResponseEntity.ok(ResponseResult.success(user));
    }

    /**
     * 创建用户
     * POST /users
     * @param user 用户对象（需要name和roleId）
     * @return 创建的用户对象（包含角色信息）
     */
    @PostMapping
    public ResponseEntity<ResponseResult<User>> createUser(@RequestBody User user) {
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
    @PutMapping("/{id}")
    public ResponseEntity<ResponseResult<User>> updateUser(
            @PathVariable Long id,
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
    @DeleteMapping("/{id}")
    public ResponseEntity<ResponseResult<Object>> deleteUser(@PathVariable Long id) {
        userService.deleteUser(id);
        return ResponseEntity.ok(ResponseResult.success("用户删除成功"));
    }
}
