package ledgerwalletsdk

import (
	"net/http"
	"time"
)

// Config SDK 配置
type Config struct {
	BaseURL        string
	ClientID       string
	ClientSecret   string
	TokenURL       string
	HTTPClient     *http.Client
	DefaultTimeout time.Duration
}