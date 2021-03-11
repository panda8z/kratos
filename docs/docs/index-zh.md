---
id: index
slug: /zh/
title: 简介
---

ORY Kratos 是根据 [云架构最佳实践](https://www.ory.sh/docs/ecosystem/software-architecture-philosophy)
构建的API优先的身份和用户管理系统。
它实现了几乎每一个软件应用程序都需要实现的核心使用功能，如下：

- **登录和注册**：允许终端用户创建和登录其账户，使用用户名/电子邮件和密码组合、社交登录（“使用Google、GitHub登录”），
  无密码免验证登录等。
- **多条件身份验证（MFA/2FA）**：支持 TOTP（[RFC 6238](https://tools.ietf.org/html/rfc6238) ，
  [IETF RFC 4226](https://tools.ietf.org/html/rfc4226) 和更为人熟知的
  [Google Authenticator](https://en.wikipedia.org/wiki/Google_Authenticator)）等协议
- **帐户验证**：验证电子邮件地址、电话号码或物理地址是否确实属于该身份。
- **帐户恢复**：使用“忘记密码”工作流、安全代码（在 MFA 设备丢失的情况下）等恢复访问。
- **用户档案和帐户管理**：更新密码，个人资料，电子邮件地址，链接的社会档案使用安全流。
- **管理员级 API**：导入、更新、删除认证信息。

认证是一个比较棘手的问题，但 ORY Kratos 给出了一个独特的解决方案。
我们最看重的价值是安全，高效，云原生技术（比如Kubernetes）友好。

- ORY Kratos 并没有 HTML 渲染引擎。你可以自己根据自己的需要，使用擅长的编程语言和框架来构建适合自己的 UI（当然，也可以借鉴或使用我们提供的示例UI）。

- **工作流引擎**让你可以完全自定义适合自己的用户体验。不管是你的用户需要注册后立即激活，还是需要多步骤的（渐进式）注册流程，ORY Kratos 都能轻松胜任。

- 一个统一**认证数据模型**是远远不够的。你完全可以简单地定制自己的数据模型，只需要你将不同的**数据模型**写进 [JSON Schema](https://json-schema.org/)，
  然后 ORY Kratos 就能正常的为你工作了。而不是反过来，需要你修改程序以生成 **JSON Schema**，然后再让系统按照意愿执行。
  这种方式让复杂的认证工作变得简单，例如一个典型的场景：你的用户可能需要一个账单地址，分配给成本中心的内部支持人员和4楼的智能冰箱。

要了解更多 ORY Kratos 和其他开源解决方案的不同和比较，请阅读[核心概念](./concepts/index.md)和[同类对比](./further-reading/comparison.md)

ORY Kratos 并不是只关心功能的实现。它还能在绝大多数处理器（i386, amd64,
arm等）支持的主流系统（Linux, macOS, Windows）上运行。此外，它编译后是一个单一的二进制可执行文件，没有任何依赖库。
这样就可以使用最简单的操作系统环境来运行它，比如，可以直接在 Linux 内核上运行 ORY Kratos。
ORY Kratos 编译后的一个完整的Docker镜像也只有20MB大小，这就非常有利于在云原生技术栈中使用它。

ORY Kratos 做到了无负担的系统规模的水平扩展。它的唯一外部依赖是 RDBMS（关系数据库管理系统 Relational Database Management System）。
我们现在已经支持了 SQLite，PostgreSQL， MySQL 和 CockroachDB。
扩展 ORY Kratos 系统规模的时候你完全不需要关心 memcached，etcd 或其他系统组件。

ORY 项目（ORY作为一个组织有多个优秀的开源项目）设计时候的一项指导原则就是**关注点分离**。
在 ORY Kratos 项目中我们也践行这一原则。 因此，我们构建的用于解决特定问题的软件，
不仅，可以很好地解决问题，并将邻近的问题(如用户界面)转移到其他应用程序。
