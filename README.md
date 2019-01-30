# tools-UICreator
本脚本目的是将coding中重复劳动取代,只要创建好相应模板,就能生成相应的代码,之后将有mac应用版本


## 使用条件
* golang编译器,安装请参考:http://www.runoob.com/go/go-environment.html
* 官网安装包地址:https://golang.google.cn/dl/
* 如果不想安装,也可以直接运行工程目录下的`start`可执行文件


## 使用方法
* 下载工程,打开终端并cd到工程所在目录,目录下包含了`coding.go`文件,和`__templates`文件夹
* `coding.go`文件为主文件
* `__templates`文件夹中存放了各种语言的模板
* 输入`go run coding.go`运行程序,根据提示进行使用:
```
bogon:tools-UICreator teiharufumi$ go run coding.go 
-----------------------------------------
-----------------------------------------
** 欢迎使用代码机器,不会偷懒的猿不是好猴子~
** 小提示:
**   查询可用模板: --s .
**   退出: --exit
-----------------------------------------
0 : oc
1 : javascript
-----------------------------------------
请从上面列表中选择一门语言类型,序号或者名字都看你心情:0
您选择的语言是:oc
-> 
```
* 系统级指令以`--`2根短线开头,模板指令以`-`1根短线开头,请注意区分
* 自定义模板指令字符无长度限制
* 输入`-g UIImageView img`,工程将按照模板将替换后的代码输出,同时会放到粘贴板上.
```
请从上面列表中选择一门语言类型,序号或者名字都看你心情:0
您选择的语言是:oc
-> -g UIImageView img
=== ↓↓↓↓↓ === 代码已经生成,并已自动复制到粘贴板中,请收好 === ↓↓↓↓↓ ====

@property (strong, nonatomic) UIImageView *img;

- (UIImageView *)img {
    if (!_img) {
        _img = [[UIImageView alloc] init];
    }
    return _img;
}



=== ↑↑↑↑↑ === 代码已经生成,并已自动复制到粘贴板中,请收好 === ↑↑↑↑↑ ===
```

## 模板创建
* 文件名
  * 文件名可以写成`g_UIButton.txt`,使用命令`-g UIButton param1,param2`来指定此模板
  * 文件名可以写成`g_.txt`,表示未找到对应name的文件时,默认时使用的模板
* 
  
  
* 响应命名类型(取文件名中`_`前的字符串)
  * #cmd
* 模板名称(取文件名中`_`后的字符串)
  * #name
* 模板输入格式(多个输入源以`,`分隔
  * #format
* 功能说明
  * #rem
* 模板内容(#content以下的所有字符都将被列入)
  * #content
* 文件名规范
  * 响应命名类型 + '_' + 模板名称,如 `g_UIImageView.txt`
* 其它说明
  * 输入格式中,如`-g UIContrl control`,第一个参数中'g'会传值给#cmd,第二个参数中'UIControl'会传值给#name,其它参数分别使用#format中定义的其它参数分别使用#format中定义的字符串变量来接收,如'control'会传值给'instance_name'
  * #format 中,如果需要多个参数,以逗号分隔,如:instance_name1,instance_name2
* 样例
```
#cmd
#name
#format
instance_name
#rem
创建一对相应类型的Getter,Setter方法
#content
@property (strong, nonatomic) <.name.> *<.instance_name.>;

- (<.name.> *)<.instance_name.> {
    if (!_<.instance_name.>) {
        _<.instance_name.> = [[<.name.> alloc] init];
    }
    return _<.instance_name.>;
}
```


## 已有模板
* 模板可以自行补充,以语言名称为文件夹
* `oc`
* `javascript`


## 示例
* 快速创建UIImageView的实例
  * 输入`-g UIImageView backImage`
  * 输出相应的代码
