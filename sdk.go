package ledgerwalletsdk

import (
	"context" // Keep context if logger methods use it
	"errors"

	// Needed for WithHTTPClient option type
	// Needed for WithTimeout option type
	// Import internal packages
	internalclient "github.com/a19ba14d/ledger-wallet-sdk/internal/client"
	internalservice "github.com/a19ba14d/ledger-wallet-sdk/internal/service"
	// Import the new log package
)

// New 使用函数式选项初始化并返回一个新的 IWallet 服务实例。
// 必须提供 WithBaseURL 选项。
func New(opts ...Option) (IWallet, error) {
	// 1. 应用选项到内部配置
	cfg := newDefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	// 2. 基础配置校验
	if cfg.baseURL == "" {
		return nil, errors.New("ledgerwalletsdk: WithBaseURL option is required")
	}

	// 3. 使用配置创建内部客户端
	// 将根配置映射到 internalclient.Config
	clientCfg := internalclient.Config{
		BaseURL:    cfg.baseURL,
		HTTPClient: cfg.httpClient, // Pass the configured http client
		// Map other relevant fields if added (e.g., auth)
	}
	internalClient, err := internalclient.NewClient(clientCfg)
	if err != nil {
		// Log the error using the configured logger (which is now log.Logger) before returning
		if cfg.logger != nil { // Check if logger is configured
			cfg.logger.Errorf(context.Background(), "Failed to create internal client: %v", err) // Use Background context for internal logging
		}
		return nil, Wrap(err, "failed to create internal client") // Wrap error using the exported Wrap function
	}

	// 4. 使用内部客户端和日志记录器创建内部服务实例
	// The logger is already configured in cfg.logger (type log.Logger)
	service := internalservice.NewService(internalClient, cfg.logger) // Pass logger to service

	// 5. 返回公共接口 IWallet (service 实现了该接口)
	if cfg.logger != nil { // Check if logger is configured
		cfg.logger.Infof(context.Background(), "ledger-wallet-sdk initialized successfully for BaseURL: %s", cfg.baseURL)
	}
	return service, nil
}
