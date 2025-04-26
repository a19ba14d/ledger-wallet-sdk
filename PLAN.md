# 规划方案：构建高质量的 `ledger-wallet-sdk` Go 模块

本计划旨在将 `ledger-wallet-sdk` 项目构建成一个高质量、易于维护和使用的 Go 模块。

## 1. 确认模块路径和可见性

*   **目标:** 将 `github.com/a19ba14d/ledger-wallet-sdk` 作为公共 Go 模块发布。
*   **动作:**
    *   保持 `go.mod` 中的模块路径不变。
    *   确保对应的 GitHub 仓库设置为公开。

## 2. 定义和清理公共 API (已优化)

*   **目标:** 提供清晰、稳定、解耦且易于配置的公共 API。
*   **动作:**
    *   **主要入口 (优化 - 函数式选项):**
        *   修改根目录 `sdk.go` 中的 `New()` 函数，采用函数式选项模式。
        *   定义 `Option` 类型和各种 `WithXxx` 函数（如 `WithBaseURL`, `WithTimeout`, `WithCredentials`, `WithLogger`）。
        *   `New()` 函数接收 `...Option` 作为参数，应用选项并返回 `WalletService` 实例。
    *   **核心接口:** 确认根目录 `service.go` 中的 `IWallet` 接口定义核心功能。
    *   **移除冗余初始化:** 移除旧 `service.go` 中的 `init()` 和 `Wallet()` 单例。
    *   **API 模型 (优化 - 解耦):**
        *   在根目录创建 `types.go` (或扩展 `models.go`)，定义 SDK 自己的、独立的输入/输出结构体（如 `Wallet`, `Balance`, `Hold`, `Monetary`, `Subject`, `Transaction` 等）。
        *   修改 `IWallet` 接口的方法签名，使用这些 SDK 自定义的类型。
    *   **输入参数:** 保持使用根目录 `models.go` 中的 `*Params` 结构体作为列表/查询操作的输入。
    *   **日志接口:** 检查根目录 `logger.go` 中的 `Logger` 接口定义。

## 3. 代码组织和封装 (已优化)

*   **目标:** 强制隐藏内部实现细节。
*   **动作:**
    *   **创建 `internal` 目录:**
        *   将 `generated/v1` 目录移动到 `internal/generated/v1`。
        *   将封装 `generated/v1` 客户端交互的逻辑（可能在 `wallet_client.go` 或 `client.go`）移入 `internal` 目录（例如 `internal/client/client.go`）。
        *   将 `IWallet` 接口的实现 (`sWallet`) 移入 `internal` 目录（例如 `internal/service/service.go`）。
        *   更新项目内的导入路径以反映这些变化。
    *   **错误处理:** 检查 `internal/service/service.go` 中的错误处理，确保返回有意义的错误。


## 4. 版本管理与发布

*   **目标:** 采用语义化版本控制 (SemVer)。
*   **动作:** 使用 Git 标签发布版本 (e.g., `v0.1.0`)，遵循 SemVer 规则。