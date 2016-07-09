## 门店标准化项目

1. 运行部署：go build, 然后运行生成的执行文件 ./acc-assign 即启动了服务。
    在浏览器输入 localhost:9000查看结果。  

2. 数据导入成js： 目前展示数据的方式是通过数据文件生成js文件。 ./acc-assign --inport-js  data-files...
  data-file在raw-data文件夹中。

3. 数据导入到数据库：也支持了raw-data到数据库的导入。目前使用的是sqlite3数据库。
   ./acc-assign --import-db data-files...
   通过restful接口向页面暴露数据。

