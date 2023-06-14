# Feishu2Md

[![Golang - feishu2md](https://img.shields.io/github/go-mod/go-version/wsine/feishu2md?color=%2376e1fe&logo=go)](https://go.dev/)
[![Unittest](https://github.com/Wsine/feishu2md/actions/workflows/unittest.yaml/badge.svg)](https://github.com/Wsine/feishu2md/actions/workflows/unittest.yaml)
[![Release](https://img.shields.io/github/v/release/wsine/feishu2md?color=orange&logo=github)](https://github.com/Wsine/feishu2md/releases)
[![Docker - feishu2md](https://img.shields.io/badge/Docker-feishu2md-2496ed?logo=docker&logoColor=white)](https://hub.docker.com/repository/docker/wwwsine/feishu2md)
[![Render - feishu2md](https://img.shields.io/badge/Render-feishu2md-4cfac9?logo=render&logoColor=white)](https://feishu2md.onrender.com)

这是一个下载飞书文档为 Markdown 文件的工具，使用 Go 语言实现。 在https://github.com/Wsine/feishu2md项目上进行的修改,自己使用顺便共享.

## 获取 API Token

配置文件需要填写 APP ID 和 APP SECRET 信息，请参考 [飞书官方文档](https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/get-) 获取。推荐设置为

- 进入飞书[开发者后台](https://open.feishu.cn/app)
- 创建企业自建应用，信息随意填写
- 选择测试企业和人员，创建测试企业，绑定应用，切换至测试版本
- （重要）打开权限管理，云文档，开通所有只读权限
  - 「查看、评论和导出文档」权限 `docs:doc:readonly`
  - 「查看 DocX 文档」权限 `docx:document:readonly`
  - 「查看、评论和下载云空间中所有文件」权限 `drive:drive:readonly`
  - 「查看和下载云空间中的文件」权限 `drive:file:readonly`
- 打开凭证与基础信息，获取 App ID 和 App Secret
### wiki专用
- 新建群组,把应用加入到群内(添加机器人选项)
- 在知识库设置里面->成员设置->把群组添加进来.

[请我喝咖啡](https://bmc.link/linkliang)

<details>
  <summary>命令行版本</summary>
  
  借助 Go 语言跨平台的特性，已编译好了主要平台的可执行文件，可以在 [Release](https://github.com/Wsine/feishu2md/releases) 中下载，并将相应平台的 feishu2md 可执行文件放置在 PATH 路径中即可。
   
   **查阅帮助文档**

   ```bash
   $ feishu2md -h
   NAME:
      feishu2md - download feishu/larksuite document to markdown file

   USAGE:
      feishu2md [global options] command [command options] [arguments...]

   VERSION:
      v2-1f5416e

   COMMANDS:
      config   Read config file or set field(s) if provided
      dump     Dump json response of the OPEN API
      help, h  Shows a list of commands or help for one command

   GLOBAL OPTIONS:
      --help, -h     show help (default: false)
      --version, -v  print the version (default: false)

   $ feishu2md config -h
   NAME:
      feishu2md config - Read config file or set field(s) if provided

   USAGE:
      feishu2md config [command options] [arguments...]

   OPTIONS:
      --appId value      Set app id for the OPEN API
      --appSecret value  Set app secret for the OPEN API
      --help, -h         show help (default: false)
   ```

   **生成配置文件**

   通过 `feishu2md config --appId <your_id> --appSecret <your_secret>` 命令即可生成该工具的配置文件。

   通过 `feishu2md config` 命令可以查看配置文件路径以及是否成功配置。

   更多的配置选项请手动打开配置文件更改。

   **下载为 Markdown**

   通过 `feishu2md <your feishu docx url>` 直接下载，文档链接可以通过 **分享 > 开启链接分享 > 复制链接** 获得。

   示例：

   ```bash
   $ feishu2md https://domain.feishu.cn/docs/docxtoken
   ```
</details>


## 感谢

- [chyroc/lark](https://github.com/chyroc/lark)
- [chyroc/lark_docs_md](https://github.com/chyroc/lark_docs_md)
- [Wsine/feishu2md](https://github.com/Wsine/feishu2md)
