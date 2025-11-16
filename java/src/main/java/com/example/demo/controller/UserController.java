package com.example.demo.controller;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/users")
public class UserController {
    // 模拟数据库
    private List<Map<String, Object>> users = new ArrayList<>();

    public UserController() {
        users.add(Map.of("id", 101, "name", "张三", "role", "学生"));
        users.add(Map.of("id", 102, "name", "李四", "role", "老师"));
    }

    @GetMapping
    public List<Map<String, Object>> getUsers() {
        return users;
    }
}
