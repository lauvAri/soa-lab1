```bash
mvn spring-boot:run
```

正常启动输出如下：

```

  .   ____          _            __ _ _
 /\\ / ___'_ __ _ _(_)_ __  __ _ \ \ \ \
( ( )\___ | '_ | '_| | '_ \/ _` | \ \ \ \
 \\/  ___)| |_)| | | | | || (_| |  ) ) ) )
  '  |____| .__|_| |_|_| |_\__, | / / / /
 =========|_|==============|___/=/_/_/_/

 :: Spring Boot ::                (v3.5.7)

2025-11-16T17:15:59.951+08:00  INFO 12716 --- [demo] [           main] com.example.demo.DemoApplication         : Starting DemoApplication using Java 21.0.5 with PID 12716 (D:\test\java\target\classes started by 19871 in D:\test\java)
2025-11-16T17:16:01.089+08:00  INFO 12716 --- [demo] [           main] o.apache.catalina.core.StandardService   : Starting service [Tomcat]
2025-11-16T17:16:01.089+08:00  INFO 12716 --- [demo] [           main] o.apache.catalina.core.StandardEngine    : Starting Servlet engine: [Apache Tomcat/10.1.48]
2025-11-16T17:16:01.191+08:00  INFO 12716 --- [demo] [           main] o.a.c.c.C.[Tomcat].[localhost].[/]       : Initializing Spring embedded WebApplicationContext
2025-11-16T17:16:01.193+08:00  INFO 12716 --- [demo] [           main] w.s.c.ServletWebServerApplicationCoe: [Apache Tomcat/10ee: [Apache Tomcat/10.1.48]
2025-11-16T17:16:01.191+08:00  INFO 12716 --- [demo] [           main] o.a.c.c.C.[Tomcat].[localhost].[/]       : Initializing Spring embedded WebApplicationContext
2025-11-16T17:16:01.193+08:00  INFO 12716 --- [demo] [           main] w.s.c.ServletWebServerApplicationContext : Root WebApplicationContext: initialization completed in 1182 ms
2025-11-16T17:16:01.645+08:00  INFO 12716 --- [demo] [           main] o.s.b.w.embedded.tomcat.TomcatWebServer  : Tomcat started on port 8080 (http) with context path '/'
2025-11-16T17:16:01.657+08:00  INFO 12716 --- [demo] [           main] com.example.demo.DemoApplication         : Started DemoApplication in 2.247 seconds (process running for 2.652)
2025-11-16T17:17:24.664+08:00  INFO 12716 --- [demo] [nio-8080-exec-1] o.a.c.c.C.[Tomcat].[localhost].[/]       : Initializing Spring DispatcherServlet 'dispatcherServlet'
```