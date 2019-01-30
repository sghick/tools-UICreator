# tools-UICreator
本脚本目的是将coding中重复劳动取代,只要创建好相应模板,就能生成相应的代码,之后将有mac应用版本


## 自带模板
* `oc`语言
  * `classes`类创建模板
  * `getter_setter`getter/setter方法创建模板


## 使用方法
* 运行脚本
* `.`表示以`__templates`路径为根目录的路径
* 输入`--s oc/classes`查询classes目录下的所有模板名,输入格式,作用
* 输入`--s .`查询脚本目录下的所有模板名,输入格式,作用
* 输入`--s . UIIma`查询指定目录下所有模板名字以`UIIma`开头的模板


## 模板创建
* 文件名
  * 文件名可以写成`g_UIButton.txt`,使用
  * 文件名可以写成`g_.txt`,表示未找到对应name的文件时,默认时使用的模板
  
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
```objc
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


## 示例
* 快速创建UIImageView的实例
  * 输入`-g UIImageView backImage`
  * 输出相应的代码
