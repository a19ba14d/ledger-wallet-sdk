# Ledger Wallet SDK 重构计划

## 目标

重构 SDK 的初始化流程，实现以下目标：

1.  **统一入口**: 通过 `sdk.go` 中的 `New` 函数提供唯一的初始化入口。
2.  **配置驱动**: 初始化过程由用户通过函数式选项传入的配置驱动。
3.  **配置校验**: 在初始化时校验必要的配置参数。
4.  **依赖注入**: 将初始化好的客户端实例注入到服务实例中。
5.  **单例模式**: 确保通过 `New` 函数获取的服务实例是全局单例。
6.  **保持服务逻辑**: 不改变服务层方法的具体业务实现。

## 重构步骤

1.  **定义配置与选项 (`options.go`)**:
    *   创建 `options.go` 文件（如果不存在）。
    *   定义 `Config` 结构体，包含：
        *   `BaseURL string` (必需)
        *   `HTTPClient *http.Client` (可选, 允许用户自定义底层 HTTP 客户端)
    *   为 `Config` 添加 `Validate()` 方法，检查 `BaseURL` 是否为空。
    *   定义函数式选项类型 `type Option func(*Config)`。
    *   实现选项函数：
        *   `func WithBaseURL(url string) Option`
        *   `func WithHTTPClient(client *http.Client) Option`

2.  **修改客户端 (`internal/client/wallet_client.go`)**:
    *   移除现有的单例逻辑 (`insWalletClient`, `walletClientOnce`, `WalletClient()` 函数)。
    *   移除从 `g.Cfg()` 读取配置的逻辑。
    *   修改 `sWalletClient` 结构体，使其直接持有初始化好的 `*walletsclient.APIClient`。
    *   创建一个新的构造函数 `func NewWalletClient(cfg *ledgerwalletsdk.Config) (IWalletClient, error)`:
        *   接收 `ledgerwalletsdk.Config` 指针作为参数。
        *   根据 `cfg.BaseURL` 和 `cfg.HTTPClient` (如果提供) 创建并配置 `walletsclient.Configuration`。
        *   **注意**: 暂时忽略 OAuth2 相关参数 (`ClientID`, `ClientSecret`, `TokenURL`) 的处理。
        *   使用配置好的 `walletsclient.Configuration` 创建 `walletsclient.APIClient`。
        *   返回包含已创建 `APIClient` 的 `*sWalletClient` 实例和 `nil` 错误，或在失败时返回 `nil` 和错误。
    *   简化 `GetClient` 方法，可能直接返回存储的 `*walletsclient.APIClient`。

3.  **修改服务 (`internal/service/service.go`)**:
    *   在 `sWallet` 结构体中添加一个字段：`client internalClient.IWalletClient` (注意导入路径可能需要调整)。
    *   修改 `NewWallet` 函数，使其接受一个 `internalClient.IWalletClient` 参数，并在创建 `sWallet` 实例时将其存储到 `client` 字段：`func NewWallet(client internalClient.IWalletClient) *sWallet`。
    *   移除 `init()` 函数、全局变量 `insWallet` 以及 `Wallet()` 函数。

4.  **实现 SDK 入口 (`sdk.go`)**:
    *   定义包级别的变量用于实现单例模式：`var sdkInstance IWallet` 和 `var sdkOnce sync.Once`。
    *   实现 `func New(opts ...Option) (IWallet, error)` 函数：
        *   使用 `sdkOnce.Do` 来确保初始化逻辑只执行一次。
        *   在 `Do` 的闭包内：
            *   创建默认的 `Config` 实例。
            *   应用所有传入的 `opts` 到 `Config`。
            *   调用 `config.Validate()` 进行校验，失败则记录并返回错误。
            *   调用 `internalClient.NewWalletClient(&config)` 创建客户端实例，处理错误。
            *   调用 `internalService.NewWallet(clientInstance)` 创建服务实例。
            *   将服务实例赋值给 `sdkInstance`。
        *   检查 `Do` 过程中是否发生错误。
        *   返回 `sdkInstance` 和错误状态。

5.  **调整服务方法实现 (例如 `internal/service/service_createwallet.go` 等)**:
    *   修改所有 `sWallet` 的方法（如 `CreateWallet`, `GetWallet` 等），使其通过 `s.client` 字段获取 `IWalletClient` 实例，然后调用 `s.client.GetClient(ctx)` (或简化后的方法) 来获取底层的 `*walletsclient.APIClient`，再执行相应的 API 调用。