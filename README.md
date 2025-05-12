# ASSCAN

<div align="center">
</br>

![Version](https://img.shields.io/badge/版本-0.1.5-blue)
![Language](https://img.shields.io/badge/语言-Go-orange)
![License](https://img.shields.io/badge/许可证-GPL-green)

**一款从fscan中提取出的端口扫描器, 以进度条形式展示扫描过程**

</div>

## 📖 项目概述

Asscan是由fscan端口扫描模块提取出的端口扫描器，更改默认扫描端口为全端口扫描，增加进度条以实时展示扫描进度，享受与fscan相同的扫描速度


## 📝 TodoList
~~1. 新增：扫描结果高亮显示~~

~~2. 新增：高频端口优先排队扫描~~

## ✨ 优势

- **🚀 快速扫描** - go原生多协程任务，纵享fscan相同的扫描速度，平均5s扫完单ip top1000端口，3min完成单ip全端口扫描
- **🔄 高频端口优先** - 将高频端口优先排队扫描，提升扫描效率
- **🔎 进度条展示** - 对扫描任务以进度条形式进行展示，实时查看任务剩余时间和任务结果
- **🔄 结果高亮** - 将扫描结果高亮展示于终端同时实时保存至文件

## 🚀 安装方法

### 预编译二进制文件

从[Releases页面](https://github.com/fxe00/asscan/releases)下载适合您系统的预编译二进制文件：

- `asscan-linux-x86_64` - Linux (64位)
- `asscan-windows-x86_64.exe` - Windows (64位)
- `asscan-macos-x86_64` - macOS (Intel)
- `asscan-macos-arm64`  - macOS (Apple Silicon)

### 从源码编译

```bash
# 克隆仓库
git clone https://github.com/fxe00/asscan.git
cd asscan

# 编译
go build -o asscan main.go

# 给予权限(windows无需)
chmod +x asscan

# 运行
./asscan -h www.baidu.com
```

### 指定参数

| 参数 | 描述 |
|------|------|
| `--help` | 显示帮助信息 |
| `-h` | (必填项)指定扫描目标host(ip/域名) |
| `-p` | 指定扫描的端口范围 (默认为1-65535) |
| `-hf` |  指定扫描目标的文件 |
| `-t` | 指定扫描的并发数 (默认为600) |
| `-timeout` | 指定端口连接的超时时间 (默认为2秒) |

## ⚙️ 连接超时设置说明

默认的端口请求超时时间为2s，这在大部分网络环境下是满足的，但该配置在网络延时较高时会出现部分端口漏报的问题。如果网络环境较差时，请使用--timeout将超时设置为3-5秒，但此操作会拉长整体的扫描时间。

## 使用示例

### 帮助
```shell
➜  asscan git:(main) ✗ go run main.go --help

   __    ___  ___   ___    __    _  _ 
  /__\  / __)/ __) / __)  /__\  ( \( )
 /(__)\ \__ \\__ \( (__  /(__)\  )  ( 
(__)(__)(___/(___/ \___)(__)(__)(_)\_)
                     asscan version: 1.0
Usage of /var/folders/wx/ndqpvv7s1hg3w5rc_zssmxyh0000gn/T/go-build340836988/b001/exe/main:
  -h string
        IP address of the host you want to scan,for example: 192.168.11.11 | 192.168.11.11-255 | 192.168.11.11,192.168.11.12
  -hf string
        host file, -hf ip.txt
  -p string
        Select a port,for example: 22 | 1-65535 | 22,80,3306 (default "1-65535")
  -t int
        Thread nums, default 600 (default 600)
  -timeout int
        Set connection timeout, default 2s (default 2)
```

### 扫描示例
```shell
➜  asscan git:(main) ✗ go run main.go -h www.baidu.com

   __    ___  ___   ___    __    _  _ 
  /__\  / __)/ __) / __)  /__\  ( \( )
 /(__)\ \__ \\__ \( (__  /(__)\  )  ( 
(__)(__)(___/(___/ \___)(__)(__)(_)\_)
                     asscan version: 1.0
start port scan
[*] www.baidu.com:80 open
[*] www.baidu.com:443 open
Scanning 7802 / 65535 [===========>-------------------------------------------------------------------------------------]  11.91% 03m22s
```

## 📜 许可证

GPL License

## 🤝 贡献

欢迎提交PR和Issue，一起改进Asscan！

---

<div align="center">
<i>更贴近日常使用的端口扫描，让生活更美好</i>
</div>