命名规则
====

### 文件名

*   整个应用或包的主入口文件应当是 `main.go` 或与应用名称简写相同。例如：`Gogs` 的主入口文件名为 `gogs.go`。

### 函数或方法

*   若函数或方法为判断类型（返回值主要为 `bool` 类型），则名称应以 `Has`, `Is`, `Can` 或 `Allow` 等判断性动词开头：
    
    ```golang

    func HasPrefix(name string, prefixes []string) bool { ... }
    func IsEntry(name string, entries []string) bool { ... }
    func CanManage(name string) bool { ... }
    func AllowGitHook() bool { ... }

    ```

### 常量

*   常量均需使用全部大写字母组成，并使用下划线分词：
    
    ```golang
    
    const APP_VER = "0.7.0.1110 Beta"
    
    ```
    
*   如果是枚举类型的常量，需要先创建相应类型：
    
    ```golang
    
    type Scheme string
    
    const (
       HTTP  Scheme = "http"
       HTTPS Scheme = "https"
    )
    
    
    ```
    
*   如果模块的功能较为复杂、常量名称容易混淆的情况下，为了更好地区分枚举类型，可以使用完整的前缀：
    
    ```golang
    
    type PullRequestStatus int

    const (
       PULL_REQUEST_STATUS_CONFLICT PullRequestStatus = iota
       PULL_REQUEST_STATUS_CHECKING
       PULL_REQUEST_STATUS_MERGEABLE
    )
    
    
    ```
    

### 变量

*   变量命名基本上遵循相应的英文表达或简写。
*   在相对简单的环境（对象数量少、针对性强）中，可以将一些名称由完整单词简写为单个字母，例如：
    *   `user` 可以简写为 `u`
    *   `userID` 可以简写 `uid`
*   若变量类型为 `bool` 类型，则名称应以 `Has`, `Is`, `Can` 或 `Allow` 开头：
    
    ```golang
    
    var isExist bool
    var hasConflict bool
    var canManage bool
    var allowGitHook bool
    
    
    ```
    
*   上条规则也适用于结构定义：
    
    ```golang
    
    // Webhook represents a web hook object.
    type Webhook struct {
        ID           int64
        RepoID       int64
        OrgID        int64
        URL          string
        ContentType  HookContentType
        Secret       string
        Events       string
        IsSSL        bool
        IsActive     bool
        HookTaskType HookTaskType
        Meta         string
        LastStatus   HookStatus
        Created      time.Time
        Updated      time.Time
    }
    
    
    ```
    

#### 变量命名惯例

变量名称一般遵循驼峰法，但遇到特有名词时，需要遵循以下规则：

*   如果变量为私有，且特有名词为首个单词，则使用小写，如 `apiClient`。
*   其它情况都应当使用该名词原有的写法，如 `APIClient`、`repoID`、`UserID`。

下面列举了一些常见的特有名词：

```golang

// A GonicMapper that contains a list of common initialisms taken from golang/lint
var LintGonicMapper = GonicMapper{
    "API":   true,
    "ASCII": true,
    "CPU":   true,
    "CSS":   true,
    "DNS":   true,
    "EOF":   true,
    "GUID":  true,
    "HTML":  true,
    "HTTP":  true,
    "HTTPS": true,
    "ID":    true,
    "IP":    true,
    "JSON":  true,
    "LHS":   true,
    "QPS":   true,
    "RAM":   true,
    "RHS":   true,
    "RPC":   true,
    "SLA":   true,
    "SMTP":  true,
    "SSH":   true,
    "TLS":   true,
    "TTL":   true,
    "UI":    true,
    "UID":   true,
    "UUID":  true,
    "URI":   true,
    "URL":   true,
    "UTF8":  true,
    "VM":    true,
    "XML":   true,
    "XSRF":  true,
    "XSS":   true,
}

```
