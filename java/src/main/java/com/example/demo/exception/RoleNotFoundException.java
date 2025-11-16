package com.example.demo.exception;

/**
 * 角色不存在异常
 */
public class RoleNotFoundException extends RuntimeException {
    public RoleNotFoundException(String message) {
        super(message);
    }

    public RoleNotFoundException(Long id) {
        super("角色不存在，ID: " + id);
    }
}

