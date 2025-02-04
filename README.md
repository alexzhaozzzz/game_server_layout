## 后端服务


## 关于协议
- 前端->后端 && 后端内部（game_server->db_server）统一使用PB协议
- 协议文件统一在协议库进行管理，包括后端内部使用协议码及协议结构
- 之前`旧逻辑`中使用自定义协议可保留，后续看情况是否重构

## 关于编译
- Windows 下直接执行 build.bat
- Linux 下查看 makefile 相关命令

## 关于代码规范
### 1、文件夹命名
- 项目根目录第一层的文件夹使用首字母大写,驼峰命名
- 其他文件夹都以小写单词，使用下划线分隔

### 2、文件命名
- 尽量采取有意义的文件名，简短，有意义，应该为小写单词，使用下划线分隔各个单词
- `msg_tile.go` 

### 3、 结构体命名
- 采用驼峰命名法，首字母根据访问控制大写或者小写

### 4、 接口命名
- 命名规则基本和上面的结构体类型
- 单个函数的结构名以 “er” 作为后缀，例如 Reader , Writer 。

### 5、变量命名
- 和结构体类似，变量名称一般遵循驼峰法，首字母根据访问控制原则大写或者小写，但遇到特有名词时，需要遵循以下规则：
- 如果变量为私有，且特有名词为首个单词，则使用小写，如 apiClient
- 其它情况都应当使用该名词原有的写法，如 APIClient、repoID、UserID
- 错误示例：UrlArray，应该写成 urlArray 或者 URLArray
- 若变量类型为 bool 类型，则名称应以 Has, Is, Can 或 Allow 开头

### 6、常量命名
- 常量均需使用全部大写字母组成，并使用下划线分词
- `const APP_VER = "1.0"`
- 如果是枚举类型的常量，需要先创建相应类型：
```
type Scheme string

const (
  HTTP  Scheme = "http"
  HTTPS Scheme = "https"
)
```

## 代码风格

### 1、缩进和折行
- 缩进直接使用 gofmt 工具格式化即可（gofmt 是使用 tab 缩进的）；
- 折行方面，一行最长不超过120个字符，超过的请使用换行展示，尽量保持格式优雅。
- 我们使用Goland开发工具，可以直接使用快捷键：ctrl+alt+L，即可。
 
### 2、语句的结尾
- Go语言中是不需要类似于Java需要冒号结尾，默认一行就是一条数据
- 如果你打算将多个语句写在同一行，它们则必须使用 ;

### 3、括号和空格
- 括号和空格方面，也可以直接使用 gofmt 工具格式化（go 会强制左大括号不换行，换行会报语法错误），所有的运算符和操作数之间要留空格。
```azure
// 正确的方式
if a > 0 {

} 

// 错误的方式
if a>0  // a ，0 和 > 之间应该空格
{       // 左大括号不可以换行，会报语法错误

}
```