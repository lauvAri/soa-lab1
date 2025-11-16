package com.example.demo.exception;

/**
 * 用户不存在异常
 */
public class UserNotFoundException extends RuntimeException {
    public UserNotFoundException(String message) {
        super(message);
    }

    public UserNotFoundException(Long id) {
        super("用户不存在，ID: " + id);
    }
}

