```shell
bee api .
```

##### beego session
[文章](http://wlb.wlb.blog.163.com/blog/static/46741320152541954128/)


##### 三目运算符
golang 思想：一种事情有且只有一种方法完成


##### beego conf
[文章](http://blog.csdn.net/qq_33610643/article/details/53511058)


##### install
```shell
go get github.com/astaxie/beego

git clone git@github.com:golang/net.git ~/gocode/src/golang.org/x/net
go get github.com/beego/admin
```


##### admin
```shell
: "复制静态文件"
cp -R $GOPATH/src/github.com/beego/admin/static ./
cp -R $GOPATH/src/github.com/beego/admin/views ./

: "初始化数据库表
会创建一个用户名、密码都是 admin 的用户
如果登录不了：清缓存
"
./securityDoor -syncdb
```


##### run
```shell
: "
-gendoc=true  每次自动化的 build 文档
-downdoc=true 自动下载 swagger 文档查看器
访问 http://***/swagger
"
bee run -gendoc=true -downdoc=true
```
