package com.example.demo.controller;

import java.util.HashMap;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;
import org.springframework.http.ResponseEntity;

@RestController
@RequestMapping("/api")
public class GatewayController {

    @Autowired
    private RestTemplate restTemplate;

    // Go 服务地址
    private final String MATERIAL_SERVICE_URL = "http://localhost:8082/materials";
    // Python 服务地址
    private final String BORROW_SERVICE_URL = "http://localhost:8081/borrows";

    // 1. 转发请求到 Go 服务 (获取物资)
    @GetMapping("/materials")
    public Object getMaterials() {
        // 远程调用 Go 接口
        ResponseEntity<Object> response = restTemplate.getForEntity(MATERIAL_SERVICE_URL, Object.class);
        return response.getBody();
    }

    // 2. 转发请求到 Python 服务 (获取借阅记录)
    @GetMapping("/borrows")
    public Object getBorrows() {
        // 远程调用 Python 接口
        ResponseEntity<Object> response = restTemplate.getForEntity(BORROW_SERVICE_URL, Object.class);
        return response.getBody();
    }

    // 3. 聚合接口示例 (同时获取人和物资 - SOA 的体现)
    @GetMapping("/dashboard")
    public Map<String, Object> getDashboard() {
        Map<String, Object> dashboard = new HashMap<>();

        // 获取本地人员数据 (简化，直接模拟)
        dashboard.put("users_count", 2);

        // 获取远程物资数据
        Object materials = restTemplate.getForObject(MATERIAL_SERVICE_URL, Object.class);
        dashboard.put("materials", materials);

        return dashboard;
    }
}
