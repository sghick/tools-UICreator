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
* 定义模板名称
  * name=UIImageView
* 定义模板输入格式
  * format=(instance_name)
* 定义功能说明
  * rem=创建一对UIImageView类型的Getter,Setter方法
* 定义模板内容
  * `content=@property (strong, nonatomic) UIImageView *<.name.>\n\n- (<.name.> *)<.instance_name.> {\n\tif (!_<.instance_name.>) {\n\t\t_<.instance_name.> = [[<.name.> alloc] init];\n\t}\n\treturn _<.instance_name.>;\n}`


## 示例
* 快速创建UIImageView的实例
  * 输入`--c UIImageView backImage`
  * 输入相应的代码
