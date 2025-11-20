package com.example.demo.common;

import java.time.LocalDateTime;

/**
 * 统一响应结果类
 */
public class ResponseResult<T> {
    /**
     * 响应码
     */
    private Integer code;

    /**
     * 响应消息
     */
    private String message;

    /**
     * 响应数据
     */
    private T data;

    /**
     * 时间戳
     */
    private LocalDateTime timestamp;

    public ResponseResult() {
        this.timestamp = LocalDateTime.now();
    }

    public ResponseResult(Integer code, String message, T data) {
        this.code = code;
        this.message = message;
        this.data = data;
        this.timestamp = LocalDateTime.now();
    }

    /**
     * 成功响应（无数据）
     */
    public static <T> ResponseResult<T> success() {
        return new ResponseResult<>(200, "success", null);
    }

    /**
     * 成功响应（有数据）
     */
    public static <T> ResponseResult<T> success(T data) {
        return new ResponseResult<>(200, "success", data);
    }

    /**
     * 成功响应（自定义消息）
     */
    public static <T> ResponseResult<T> success(String message, T data) {
        return new ResponseResult<>(200, message, data);
    }

    /**
     * 失败响应
     */
    public static <T> ResponseResult<T> error(Integer code, String message) {
        return new ResponseResult<>(code, message, null);
    }

    /**
     * 失败响应（400）
     */
    public static <T> ResponseResult<T> badRequest(String message) {
        return new ResponseResult<>(400, message, null);
    }

    /**
     * 失败响应（404）
     */
    public static <T> ResponseResult<T> notFound(String message) {
        return new ResponseResult<>(404, message, null);
    }

    /**
     * 失败响应（500）
     */
    public static <T> ResponseResult<T> serverError(String message) {
        return new ResponseResult<>(500, message, null);
    }

    // Getter和Setter方法
    public Integer getCode() {
        return code;
    }

    public void setCode(Integer code) {
        this.code = code;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public T getData() {
        return data;
    }

    public void setData(T data) {
        this.data = data;
    }

    public LocalDateTime getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(LocalDateTime timestamp) {
        this.timestamp = timestamp;
    }
}

