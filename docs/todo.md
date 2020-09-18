## feature/03/addSQlDataSource

+ 将数据类型做成枚举，替代原先的IsFlux/IsSQL判断
+ 数据源管理模块增加SQL数据库连接管理
+ 定义查询参数的模板
+ 增加SQLMaker编辑SQL
+ TimeSeries中增加执行Sql返回timeSeriesFlux数据结构的函数
+ 返回数据后，检验目前的展示组件能否配合显示

可以先尝试执行SQL返回数据的部分，验证展示是否合适
