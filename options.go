package ledgerwalletsdk

import (
	"net/http"
	"time"
)

import "github.com/a19ba14d/ledger-wallet-sdk/log" // Import the new log package

// config 应用函数式选项后的内部配置结构。
type config struct {
	baseURL        string
	defaultTimeout time.Duration
	httpClient     *http.Client
	logger         log.Logger // Use log.Logger
	// Add other options like auth credentials if needed
}

// Option 是一个函数，用于修改内部配置。
type Option func(*config)

// newDefaultConfig 创建一个带有默认值的配置。
func newDefaultConfig() *config {
	return &config{
		defaultTimeout: 15 * time.Second, // Default timeout
		logger:         log.NoopLogger{}, // Use log.NoopLogger
		// Initialize other defaults if necessary
	}
}

// WithBaseURL 设置 API 的基础 URL。
func WithBaseURL(url string) Option {
	return func(c *config) {
		c.baseURL = url
	}
}

// WithTimeout 设置默认的请求超时时间。
func WithTimeout(timeout time.Duration) Option {
	return func(c *config) {
		if timeout > 0 {
			c.defaultTimeout = timeout
		}
	}
}

// WithHTTPClient 允许提供自定义的 http.Client。
// 这对于添加自定义中间件（如认证、重试）很有用。
func WithHTTPClient(client *http.Client) Option {
	return func(c *config) {
		if client != nil {
			c.httpClient = client
		}
	}
}

// WithLogger 设置用于 SDK 内部日志记录的记录器。
func WithLogger(logger log.Logger) Option { // Use log.Logger
	return func(c *config) {
		if logger != nil {
			c.logger = logger
		}
	}
}

// TODO: Add options for authentication if needed
// func WithClientCredentials(clientID, clientSecret, tokenURL string) Option { ... }