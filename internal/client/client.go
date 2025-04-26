package client // Changed package name

import (
	"context"
	// "fmt" // Removed unused import
	"net/http"
	// "sync" // Removed unused import
	"time"

	// Updated import path for generated client
	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	// "golang.org/x/oauth2" // Commented out as OAuth2 is bypassed
	// "golang.org/x/oauth2/clientcredentials" // Commented out as OAuth2 is bypassed
)

// Client manages communication with the underlying wallets API.
// It holds the configuration and the generated API client.
type Client struct {
	generatedClient *walletsclient.APIClient
	httpClient      *http.Client // The HTTP client used for requests
	baseURL         string
	// Add other necessary config fields if needed later (e.g., auth details)
}

// Config holds the configuration needed to create an internal client.
// This should be populated by the root New function based on user-provided options.
type Config struct {
	BaseURL    string
	HTTPClient *http.Client // Allow providing a custom http client (e.g., with auth middleware)
	// Add other fields like ClientID, ClientSecret, TokenURL if OAuth2 is re-enabled
}

// NewClient creates a new internal client instance.
func NewClient(cfg Config) (*Client, error) {
	if cfg.BaseURL == "" {
		return nil, gerror.New("BaseURL is required for internal client configuration")
	}

	// Use provided HTTP client or create a default one
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		// TODO: Consider if OAuth2 should be handled here or in the root New function
		// For now, using a default client as per the original commented-out logic
		g.Log().Warning(context.Background(), "Internal Wallet API client is configured WITHOUT custom HTTP client (and potentially without OAuth2).")
		httpClient = &http.Client{Timeout: 15 * time.Second} // Default timeout
	}

	// Configure the generated walletsclient
	walletsCfg := walletsclient.NewConfiguration()
	walletsCfg.HTTPClient = httpClient
	walletsCfg.Servers = walletsclient.ServerConfigurations{
		walletsclient.ServerConfiguration{
			URL:         cfg.BaseURL,
			Description: "Default Server",
		},
	}
	// walletsCfg.Debug = ... // Optional: Add debug config if needed

	generatedClient := walletsclient.NewAPIClient(walletsCfg)

	internalClient := &Client{
		generatedClient: generatedClient,
		httpClient:      httpClient,
		baseURL:         cfg.BaseURL,
	}

	g.Log().Infof(context.Background(), "Internal Wallet API client initialized for base URL: %s", cfg.BaseURL)
	return internalClient, nil
}

// GetGeneratedClient returns the underlying generated API client.
// This is used by the internal service implementation to make API calls.
func (c *Client) GetGeneratedClient() *walletsclient.APIClient {
	// TODO: Add locking (mutex) if the client can be reconfigured concurrently,
	// but typically it's configured once on creation.
	return c.generatedClient
}

// Note: Removed old singleton logic and direct config loading.
// Configuration is now passed via the Config struct.
