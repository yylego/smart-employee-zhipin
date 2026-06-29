[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/smart-employee-zhipin/release.yml?branch=main&label=BUILD)](https://github.com/yylego/smart-employee-zhipin/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/smart-employee-zhipin)](https://pkg.go.dev/github.com/yylego/smart-employee-zhipin)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/smart-employee-zhipin/main.svg)](https://coveralls.io/github/yylego/smart-employee-zhipin?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.26+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/badge/release-active-blue.svg)](https://github.com/yylego/smart-employee-zhipin)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/smart-employee-zhipin)](https://goreportcard.com/report/github.com/yylego/smart-employee-zhipin)

# smart-employee-zhipin

**Smart Job Seeking Assistant** — A job tracking system built on the Kratos microservice framework for recording and managing the job seeking process on Boss Zhipin.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Overview

Upgrade the old Markdown-based job tracking workflow into a structured API service + admin dashboard, supporting:

- Position recording and status management (pending, messaged, interviewing, etc.)
- Requirement match analysis (compare each job requirement against resume)
- Chat sync — push entire conversation at once, auto-extract last contact time and resume status
- Firm blacklist management
- Smart filtering: stale positions, need-response, need-resend
- Admin dashboard to browse positions and chat histories

## Project Structure

```
smart-employee-zhipin/
├── zhipin-kratos/       # Kratos backend service
│   ├── api/zhipin/      # AI-facing Proto API (position/communication/match/blacklist)
│   ├── api/admin/       # Admin dashboard Proto API
│   ├── cmd/             # Entry point + Wire dependency injection
│   ├── configs/         # Configuration files
│   └── internal/        # Business logic (biz/data/service/server/enums/models)
├── zhipin-vue3/         # Vue3 + Element Plus admin dashboard frontend
├── zhipin-migrate/      # Database migration tool (golang-migrate)
├── zhipin-codegen/      # Proto → TypeScript client code generation
├── Makefile             # Top-level build entry
└── smart_employee_zhipin.go  # Root package, provides SourceRoot()
```

## Tech Stack

| Layer | Technology |
|-------|------------|
| Backend | [Kratos](https://github.com/go-kratos/kratos) v3 |
| API | Protocol Buffers + gRPC / HTTP |
| ORM | GORM + [gormcnm](https://github.com/yylego/gormcnm) (type-safe column names) |
| DI | [Wire](https://github.com/google/wire) |
| Logging | [zap](https://github.com/uber-go/zap) + [zapkratos](https://github.com/yylego/kratos-zap) |
| Frontend | Vue 3 + Element Plus + @protobuf-ts |
| Migration | [golang-migrate](https://github.com/golang-migrate/migrate) |

## Quick Start

### Setup Database

The project uses PostgreSQL via Docker:

```bash
# Create and start PostgreSQL container
docker run -d --name=postgres -e POSTGRES_PASSWORD=123 -p 55432:5432 postgres
docker update --restart=always postgres

# Create database
docker exec $(docker ps -qf "name=postgres") psql -U postgres -c "CREATE DATABASE zhipin_db;"
```

Database connection is configured in `zhipin-kratos/configs/config.yaml`.

### Generate Proto Code

Regenerate the backend proto code (Go) and auto-sync service stubs — this runs `zhipin-kratos` buf generation plus the orzkratos service sync:

```bash
make orz
```

### Generate TypeScript Clients

One command drives the whole frontend client workflow through the `zhipin-codegen` program: clean old output, buf-generate the gRPC TypeScript client, convert it into an HTTP client, sync into `zhipin-vue3/src/rpc/zhipin`, then clean up. Run it when a proto changes:

```bash
make gen
```

### Build Backend

```bash
cd zhipin-kratos
make build
```

### Start Service

```bash
cd zhipin-kratos
./bin/zhipin-kratos -conf ./configs/
```

Default listeners:
- HTTP: `0.0.0.0:8001`
- gRPC: `0.0.0.0:9001`

### Database Migration

```bash
cd zhipin-migrate
make MIGRATE-ALL
```

### Frontend Development

```bash
cd zhipin-vue3
npm install
npm run dev
```

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yylego/smart-employee-zhipin.svg?variant=adaptive)](https://starchart.cc/yylego/smart-employee-zhipin)
