package com.example.demo.controller;

import java.util.HashMap;
import java.util.Map;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.reactive.function.client.WebClient;
import reactor.core.publisher.Mono;

/**
 * Gateway 控制器
 * 注意：Spring Cloud Gateway 使用响应式编程（WebFlux）
 * 大部分路由已经在 application.yml 中配置，这里只处理聚合接口
 */
@RestController
@RequestMapping("/api")
public class GatewayController {

    private final WebClient webClient;

    // 服务地址配置
    @Value("${gateway.service.material:http://localhost:8082}")
    private String MATERIAL_SERVICE_BASE;

    @Value("${gateway.service.user:http://localhost:8083}")
    private String USER_SERVICE_BASE;

    public GatewayController(WebClient.Builder webClientBuilder) {
        this.webClient = webClientBuilder.build();
    }

    /**
     * 聚合接口示例 (同时获取用户和物资 - SOA 的体现)
     * 这个接口需要聚合多个服务的数据，所以用代码实现
     * 其他简单转发已经通过 application.yml 配置
     */
    @GetMapping("/dashboard")
    public Mono<Map<String, Object>> getDashboard() {
        // 并行调用多个服务
        Mono<Object> usersMono = webClient.get()
                .uri(USER_SERVICE_BASE + "/users")
                .retrieve()
                .bodyToMono(Object.class);

        Mono<Object> materialsMono = webClient.get()
                .uri(MATERIAL_SERVICE_BASE + "/materials")
                .retrieve()
                .bodyToMono(Object.class);

        // 等待所有服务响应后聚合
        return Mono.zip(usersMono, materialsMono)
                .map(tuple -> {
                    Map<String, Object> dashboard = new HashMap<>();
                    dashboard.put("users", tuple.getT1());
                    dashboard.put("materials", tuple.getT2());
                    dashboard.put("status", "success");
                    return dashboard;
                });
    }
}
