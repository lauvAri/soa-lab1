# User Service

## 启动方式

```bash
cd java
./mvnw clean spring-boot:run -U
```

- 必须在 `java` 目录下运行，否则 Maven 找不到 `spring-boot` 插件。
- 使用 `-U` 可以强制更新依赖，确保使用当前项目锁定的 Spring Boot `3.3.5` 与 SpringDoc 版本，避免出现 `ControllerAdviceBean.<init>(Object)` 的兼容性异常。
- 应用默认运行在 `http://localhost:8083`，数据库连接、端口等可在 `src/main/resources/application.yml` 中调整。

## Swagger / OpenAPI 文档

应用集成了 SpringDoc OpenAPI，启动后即可在浏览器中访问：

- Swagger UI：<http://localhost:8083/swagger-ui/index.html>
- OpenAPI JSON：<http://localhost:8083/v3/api-docs>

若 Swagger UI 出现 `Fetch error`，通常是 `/v3/api-docs` 返回 500。清理本地 Maven 缓存中旧的 Spring Framework 6.2 依赖（例如 `~/.m2/repository/org/springframework`），或执行 `./mvnw dependency:purge-local-repository`，重新启动即可恢复。
