package sdkconfig // Changed package name

import (
	"net/http"
	"time"
)

import (
	"fmt"
	// Updated import path for the log package
	"github.com/a19ba14d/ledger-wallet-sdk/pkg/log"
)

// Config 应用函数式选项后的配置结构。
// 这个结构体会被传递给内部包（如 client）使用。
type Config struct {
	BaseURL        string        // API 基础 URL (必需)
	HTTPClient     *http.Client  // 可选的自定义 HTTP 客户端
	defaultTimeout time.Duration // 内部使用的默认超时
	logger         log.Logger    // 内部使用的日志记录器
}

// Validate 检查配置是否有效。
func (c *Config) Validate() error {
	if c.BaseURL == "" {
		return fmt.Errorf("BaseURL is required")
	}
	// Add other validation rules if needed
	return nil
}

// Option 是一个函数，用于修改配置。
type Option func(*Config)

// NewDefaultConfig 创建一个带有默认值的配置。
// This function needs to be exported if it's used outside this package (e.g., in sdk.go)
func NewDefaultConfig() *Config {
	return &Config{
		defaultTimeout: 15 * time.Second, // Default timeout - use lowercase
		logger:         log.NoopLogger{}, // Use log.NoopLogger - use lowercase
		// Initialize other defaults if necessary
	}
}

// WithBaseURL 设置 API 的基础 URL。
func WithBaseURL(url string) Option {
	return func(c *Config) {
		c.BaseURL = url
	}
}

// WithTimeout 设置默认的请求超时时间 (内部使用)。
func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		if timeout > 0 {
			c.defaultTimeout = timeout // use lowercase
		}
	}
}

// WithHTTPClient 允许提供自定义的 http.Client。
// 这对于添加自定义中间件（如认证、重试）很有用。
func WithHTTPClient(client *http.Client) Option {
	return func(c *Config) {
		if client != nil {
			c.HTTPClient = client
		}
	}
}

// WithLogger 设置用于 SDK 内部日志记录的记录器 (内部使用)。
func WithLogger(logger log.Logger) Option { // Use log.Logger
	return func(c *Config) {
		if logger != nil {
			c.logger = logger // use lowercase
		}
	}
}

// 注意：根据计划，暂时忽略 OAuth2 相关选项。
// // func WithClientCredentials(clientID, clientSecret, tokenURL string) Option { ... }

// GetLogger returns the internal logger (needs to be exported if used externally)
// func (c *Config) GetLogger() log.Logger {
// 	return c.logger
// }

// GetDefaultTimeout returns the internal default timeout (needs to be exported if used externally)
// func (c *Config) GetDefaultTimeout() time.Duration {
// 	return c.defaultTimeout
// }