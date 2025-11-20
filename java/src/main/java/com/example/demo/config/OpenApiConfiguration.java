package com.example.demo.config;

import org.springframework.context.annotation.Configuration;

import io.swagger.v3.oas.annotations.OpenAPIDefinition;
import io.swagger.v3.oas.annotations.info.Contact;
import io.swagger.v3.oas.annotations.info.Info;
import io.swagger.v3.oas.annotations.info.License;

/**
 * Provides basic metadata for the generated OpenAPI specification.
 */
@Configuration
@OpenAPIDefinition(
		info = @Info(
				title = "用户服务接口文档",
				version = "v1",
				description = "提供用户基础资料查询、角色关联、增删改查等接口说明，方便前后端联调。",
				contact = @Contact(name = "用户服务团队"),
				license = @License(name = "Apache 2.0", url = "https://www.apache.org/licenses/LICENSE-2.0.html")
		)
)
public class OpenApiConfiguration {
}
