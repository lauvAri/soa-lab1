运行项目

```bash
go run .
```

可以看到如下输出：

```
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /materials                --> main.main.func1 (3 handlers)
[GIN-debug] POST   /materials                --> main.main.func2 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://github.com/gin-gonic/gin/blob/master/docs/doc.md#dont-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8082
```

## Swagger 文档

如果直接执行 `swag init` 会因为（只读 HOME、git safe.directory 等）环境限制而失败。可以使用仓库提供的脚本来生成/更新接口文档：

```bash
cd go
# 需要先安装 swag CLI（一次性操作）
go install github.com/swaggo/swag/cmd/swag@latest
# 使用脚本生成 docs/
./scripts/gen_swagger.sh
```

该脚本会自动将 `GOCACHE` 指到 `/tmp/go-cache` 并强制 `GOFLAGS=-buildvcs=false`，避免 `go list` 在生成过程中因为缓存或 git 权限问题而退出。
