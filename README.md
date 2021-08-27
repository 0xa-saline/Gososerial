# Gososerial

## 介绍

- ysoserial是java反序列化安全方面著名的工具
  
- 无需java环境，无需下载ysoserial.jar文件
  
- 输入命令直接获得payload，方便编写安全工具 

## Introduce

- Ysoserial is a well-known tool for Java deserialization security

- No Java environment and no need to download ysoserial.jar file

- Enter the command to directly obtain the payload, which is convenient for writing security tools

## Example

CommonsCollections1 Payload

![](https://github.com/EmYiQing/Gososerial/blob/master/img/1.png)


List of Supported

![](https://github.com/EmYiQing/Gososerial/blob/master/img/2.png)

## Quick Start

```shell
go get github.com/EmYiQing/Gososerial
```

```go
package main

import gososerial "github.com/EmYiQing/Gososerial"

func main()  {
	payload := gososerial.GetCC1("calc.exe")
	......
	sendPayload(payload)
	......
}
```

## About

参考了xray中p师傅的代码

[xray](https://xray.cool/)

[phith0n](https://github.com/phith0n)
