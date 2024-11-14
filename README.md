# examples 开源项目示例

[开源数据库 openGauss 驱动库](https://gitee.com/opengauss/openGauss-connector-go-pq)
[华为云数据库 GaussDB 驱动库](https://support.huaweicloud.com/centralized-devg-v8-gaussdb/gaussdb-42-1836.html) 该驱动为了兼容 ORM 框架，在开源版驱动的基础上，在创建数据库连接时多注册了“postgres”和“postgresql”这两个驱动名称。

- gf4opengauss：GoFrame 操作华为云 GaussDB 数据库的微服务示例，ORM 层使用 [gaussdb](https://github.com/okyer/gf/tree/master/contrib/drivers/gaussdb/v2)，底层使用 openGauss 开源的驱动库

- gorm4opengauss：GORM 操作华为云 GaussDB 数据库的微服务示例，ORM 层使用 [gorm4gaussdb](https://github.com/okyer/gorm4gaussdb)，底层使用 openGauss 开源的驱动库

- gorm4gaussdb：GORM 操作华为云 GaussDB 数据库的微服务示例，ORM 层使用 [postgres](gorm.io/driver/postgres)，底层使用华为云 GaussDB 下载的驱动库




