# Go-Printer
## 项目介绍
用于 CQUT-iCode 实验内打印机服务
- 基于：[CUPS-Github](https://github.com/adrianfalleiro/go-cups)
- CGo代码参考：[adrianfalleiro/go-cups](https://github.com/adrianfalleiro/go-cups) 部分参考，为调用 C lib 服务。

基于 CGO 调用 CUPS 函数实现打印机服务

在我本机找到的 cups.h 依赖库文件，仅会在 OSX 上进行编译，其他 Linux 发行版本可能需要手动加载依赖文件。
``` shell
/Library/Developer/CommandLineTools/SDKs/MacOSX11.1.sdk/usr/include/cups/cups.h
/Library/Developer/CommandLineTools/SDKs/MacOSX10.15.sdk/usr/include/cups/cups.h
/System/Volumes/Data/Library/Developer/CommandLineTools/SDKs/MacOSX11.1.sdk/usr/include/cups/cups.h
/System/Volumes/Data/Library/Developer/CommandLineTools/SDKs/MacOSX10.15.sdk/usr/include/cups/cups.h
```