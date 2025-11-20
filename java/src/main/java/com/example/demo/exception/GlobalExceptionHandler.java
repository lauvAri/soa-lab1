package com.example.demo.exception;

import com.example.demo.common.ResponseResult;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.http.converter.HttpMessageNotReadableException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

/**
 * 全局异常处理器
 */
@RestControllerAdvice
public class GlobalExceptionHandler {

    private static final Logger logger = LoggerFactory.getLogger(GlobalExceptionHandler.class);

    /**
     * 处理用户不存在异常
     */
    @ExceptionHandler(UserNotFoundException.class)
    public ResponseEntity<ResponseResult<Object>> handleUserNotFoundException(UserNotFoundException e) {
        ResponseResult<Object> result = ResponseResult.notFound(e.getMessage());
        return ResponseEntity.status(HttpStatus.NOT_FOUND).body(result);
    }

    /**
     * 处理角色不存在异常
     */
    @ExceptionHandler(RoleNotFoundException.class)
    public ResponseEntity<ResponseResult<Object>> handleRoleNotFoundException(RoleNotFoundException e) {
        ResponseResult<Object> result = ResponseResult.badRequest(e.getMessage());
        return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(result);
    }

    /**
     * 处理参数验证异常
     */
    @ExceptionHandler(IllegalArgumentException.class)
    public ResponseEntity<ResponseResult<Object>> handleIllegalArgumentException(IllegalArgumentException e) {
        ResponseResult<Object> result = ResponseResult.badRequest(e.getMessage());
        return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(result);
    }

    /**
     * 处理请求体缺失或格式错误异常
     */
    @ExceptionHandler(HttpMessageNotReadableException.class)
    public ResponseEntity<ResponseResult<Object>> handleHttpMessageNotReadableException(HttpMessageNotReadableException e) {
        String errorMessage = e.getMessage();
        if (errorMessage != null && errorMessage.contains("Required request body is missing")) {
            errorMessage = "请求体缺失，请确保在请求体中包含JSON格式的用户数据（name和roleId字段）";
        } else if (errorMessage != null && errorMessage.contains("JSON parse error")) {
            errorMessage = "JSON格式错误，请检查请求体的JSON格式是否正确";
        } else {
            errorMessage = "请求体格式错误: " + (errorMessage != null ? errorMessage : "未知错误");
        }
        
        ResponseResult<Object> result = ResponseResult.badRequest(errorMessage);
        return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(result);
    }

    /**
     * 处理其他异常
     */
    @ExceptionHandler(Exception.class)
    public ResponseEntity<ResponseResult<Object>> handleException(Exception e) {
        // 打印异常堆栈信息到日志，便于调试
        logger.error("服务器内部错误", e);
        
        // 获取异常消息，如果为null则使用异常类名
        String errorMessage = e.getMessage();
        if (errorMessage == null || errorMessage.trim().isEmpty()) {
            errorMessage = e.getClass().getSimpleName();
        }
        
        ResponseResult<Object> result = ResponseResult.serverError("服务器内部错误: " + errorMessage);
        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(result);
    }
}

