package client // Changed package name

import (
	"context"
	"net/http"
	"time"

	// Updated import path for sdk config
	sdkconfig "github.com/a19ba14d/ledger-wallet-sdk/pkg/sdkconfig"
	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1" // Keep generated client
	// Removed unused imports: fmt, sync, gerror, g
)

// IWalletClient defines the interface for the wallet API client service.
type IWalletClient interface {
	GetClient(ctx context.Context) (*walletsclient.APIClient, error)
	// Close can be added if explicit connection closing is needed
	// Close() error
}

// sWalletClient implements the IWalletClient interface.
// It now holds only the configured API client.
type sWalletClient struct {
	client *walletsclient.APIClient
}

// NewWalletClient creates a new wallet client instance based on the provided SDK configuration.
func NewWalletClient(cfg *sdkconfig.Config) (IWalletClient, error) { // Updated parameter type
	// Use provided http client or create a default one
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		// Note: Using the internal defaultTimeout from Config is tricky here
		// as it's not exported. We might need to rethink how timeout is handled
		// or just use a standard default here. Using 15 seconds for now.
		httpClient = &http.Client{Timeout: 15 * time.Second}
		// Consider logging a warning if using default client? Requires logger access.
	}

	// Configure the walletsclient using the provided BaseURL and httpClient
	walletsCfg := walletsclient.NewConfiguration()
	walletsCfg.HTTPClient = httpClient
	walletsCfg.Servers = walletsclient.ServerConfigurations{
		walletsclient.ServerConfiguration{
			URL:         cfg.BaseURL, // Use BaseURL from the passed Config
			Description: "Default Server",
		},
	}
	// Optional: Set Debug flag based on some config if needed later
	// walletsCfg.Debug = cfg.Debug // Assuming Debug field exists in ledgerwalletsdk.Config

	apiClient := walletsclient.NewAPIClient(walletsCfg)

	// Consider logging initialization success? Requires logger access.
	// cfg.logger.Infof(context.Background(), "Wallet API client initialized for base URL: %s", cfg.BaseURL)

	return &sWalletClient{client: apiClient}, nil
}

// GetClient returns the underlying generated API client.
// The client is now initialized eagerly in NewWalletClient.
func (s *sWalletClient) GetClient(ctx context.Context) (*walletsclient.APIClient, error) {
	// Context is passed but not used here currently. Keep it for interface consistency.
	return s.client, nil
}

// Note: Singleton logic (insWalletClient, walletClientOnce, WalletClient func) removed.
// Note: Config reading from g.Cfg() removed.
// Note: OAuth2 related code was already commented out and logic removed.
