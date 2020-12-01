### ggit


这是一个小工具，给git包装了一层，使得git clone时，使用github.com.cnpmjs.org域名替换github.com域名，从而加速clone的速度。

### Why use ggit
在国内直接从github.com clone时，速度会很慢，项目小时可能也能忍受，但项目大时，往往需要非常长的时间才能完成clone。
这时如果换成从国内的镜像站点clone，速度会非常快。

### User Guide
直接用ggit替代git命令使用即可，ggit只对clone时的url做域名替换，其它命令会直接调用git执行。

