[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/smart-employee-zhipin/release.yml?branch=main&label=BUILD)](https://github.com/yylego/smart-employee-zhipin/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/smart-employee-zhipin)](https://pkg.go.dev/github.com/yylego/smart-employee-zhipin)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/smart-employee-zhipin/main.svg)](https://coveralls.io/github/yylego/smart-employee-zhipin?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.26+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/badge/release-active-blue.svg)](https://github.com/yylego/smart-employee-zhipin)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/smart-employee-zhipin)](https://goreportcard.com/report/github.com/yylego/smart-employee-zhipin)

# smart-employee-zhipin

**智能求职助手** — 基于 Kratos 微服务框架的求职管理系统，用于记录和跟踪 Boss 直聘上的求职过程。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)

<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 项目简介

将手工 Markdown 记录的求职流程，升级为结构化的 API 服务 + 可视化管理面板，支持：

- 岗位信息录入与状态管理（待处理、已发消息、面试中等）
- 岗位需求匹配分析（逐条对比岗位要求与简历匹配度）
- 聊天记录同步——一次推送整个对话，自动提取最后沟通时间和简历状态
- 公司黑名单管理
- 智能筛选：待跟进/待回复/待补发简历
- 管理面板查看岗位和聊天记录

## 项目结构

```
smart-employee-zhipin/
├── zhipin-kratos/       # Kratos 后端服务
│   ├── api/zhipin/      # AI 调用的 Proto 接口（岗位/沟通/匹配/黑名单）
│   ├── api/admin/       # 管理面板 Proto 接口
│   ├── cmd/             # 服务入口 + Wire 依赖注入
│   ├── configs/         # 配置文件
│   └── internal/        # 业务逻辑（biz/data/service/server/enums/models）
├── zhipin-vue3/         # Vue3 + Element Plus 管理面板前端
├── zhipin-migrate/      # 数据库迁移工具（golang-migrate）
├── zhipin-codegen/      # Proto → TypeScript 客户端代码生成
├── Makefile             # 顶层构建入口
└── smart_employee_zhipin.go  # 根包，提供 SourceRoot()
```

## 技术栈

| 层级 | 技术 |
|------|------|
| 后端框架 | [Kratos](https://github.com/go-kratos/kratos) v2 |
| API 定义 | Protocol Buffers + gRPC / HTTP |
| 数据库 ORM | GORM + [gormcnm](https://github.com/yylego/gormcnm)（类型安全列名） |
| 依赖注入 | [Wire](https://github.com/google/wire) |
| 日志 | [zap](https://github.com/uber-go/zap) + [zapkratos](https://github.com/yylego/kratos-zap) |
| 前端 | Vue 3 + Element Plus + @protobuf-ts |
| 数据库迁移 | [golang-migrate](https://github.com/golang-migrate/migrate) |

## 快速开始

### 配置数据库

项目使用 PostgreSQL，通过 Docker 运行：

```bash
# 创建并启动 PostgreSQL 容器
docker run -d --name=postgres -e POSTGRES_PASSWORD=123 -p 55432:5432 postgres
docker update --restart=always postgres

# 创建数据库
docker exec $(docker ps -qf "name=postgres") psql -U postgres -c "CREATE DATABASE zhipin_db;"
```

数据库连接配置在 `zhipin-kratos/configs/config.yaml` 中。

### 生成 Proto 代码

```bash
make orz
```

### 生成 TypeScript 客户端

```bash
make gen
```

### 构建后端

```bash
cd zhipin-kratos
make build
```

### 启动服务

```bash
cd zhipin-kratos
./bin/zhipin-kratos -conf ./configs/
```

服务默认监听：
- HTTP: `0.0.0.0:8001`
- gRPC: `0.0.0.0:9001`

### 数据库迁移

```bash
cd zhipin-migrate
make MIGRATE-ALL
```

### 前端开发

```bash
cd zhipin-vue3
npm install
npm run dev
```

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->
