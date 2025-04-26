package service

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1" // 确认路径

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	// "golang.org/x/oauth2" // Commented out as OAuth2 is bypassed
	// "golang.org/x/oauth2/clientcredentials" // Commented out as OAuth2 is bypassed
)

// IWalletClient defines the interface for the wallet API client service.
type IWalletClient interface {
	GetClient(ctx context.Context) (*walletsclient.APIClient, error)
	// Close can be added if explicit connection closing is needed
	// Close() error
}

// sWalletClient implements the IWalletClient interface.
type sWalletClient struct {
	client       *walletsclient.APIClient
	baseURL      string
	clientID     string
	clientSecret string
	tokenURL     string
	httpClient   *http.Client // Store the OAuth2 authenticated client
	mutex        sync.RWMutex
}

var (
	insWalletClient  IWalletClient
	walletClientOnce sync.Once
)

// WalletClient returns an instance of the wallet client service.
func WalletClient() IWalletClient {
	walletClientOnce.Do(func() {
		insWalletClient = &sWalletClient{}
		// Lazy initialization in GetClient
	})
	return insWalletClient
}

// GetClient returns or initializes the wallet API client.
func (s *sWalletClient) GetClient(ctx context.Context) (*walletsclient.APIClient, error) {
	s.mutex.RLock()
	// Check if client and http client are already initialized
	if s.client != nil && s.httpClient != nil {
		// Note: The oauth2 client should handle token refreshing automatically.
		// If more complex checks are needed (e.g., force refresh), add logic here.
		defer s.mutex.RUnlock()
		return s.client, nil
	}
	s.mutex.RUnlock()

	// Need to create or refresh the client
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Double check in case another goroutine initialized it while waiting for the lock
	if s.client != nil && s.httpClient != nil {
		return s.client, nil
	}

	// Load configuration from manifest/config/config.yaml
	// Ensure the keys match exactly what's in the config file.
	cfgMapVar, err := g.Cfg().Get(ctx, "walletsApi")
	if err != nil {
		g.Log().Errorf(ctx, "Failed to get walletsApi configuration: %v", err)
		return nil, gerror.Wrap(err, "failed to get walletsApi configuration")
	}
	if cfgMapVar.IsNil() || !cfgMapVar.IsMap() {
		err := fmt.Errorf("walletsApi configuration is missing or not a map in manifest/config/config.yaml")
		g.Log().Error(ctx, err)
		return nil, err
	}
	cfgMap := cfgMapVar.Map()

	baseURL, okBaseURL := cfgMap["baseUrl"].(string)
	clientID, okClientID := cfgMap["oauthClientId"].(string)
	clientSecret, okClientSecret := cfgMap["oauthClientSecret"].(string)
	tokenURL, okTokenURL := cfgMap["oauthTokenUrl"].(string)

	if !okBaseURL || baseURL == "" || !okClientID || clientID == "" || !okClientSecret || clientSecret == "" || !okTokenURL || tokenURL == "" {
		err := fmt.Errorf("wallets API configuration missing or invalid in manifest/config/config.yaml (baseUrl, oauthClientId, oauthClientSecret, oauthTokenUrl must be non-empty strings)")
		g.Log().Error(ctx, err)
		return nil, err
	}

	s.baseURL = baseURL
	s.clientID = clientID
	s.clientSecret = clientSecret
	s.tokenURL = tokenURL

	// --- BEGIN COMMENTED OUT OAuth2 CODE ---
	// // Configure OAuth2 client credentials flow
	// conf := &clientcredentials.Config{
	// 	ClientID:     s.clientID,
	// 	ClientSecret: s.clientSecret,
	// 	TokenURL:     s.tokenURL,
	// 	Scopes:       []string{}, // Add required scopes if any, e.g., ["wallets:read", "wallets:write"] based on openapi.yaml
	// 	AuthStyle: oauth2.AuthStyleInParams, // Adjust if needed, default is usually AuthStyleInHeader
	// }
	//
	// // Create an HTTP client that automatically handles token refresh
	// // Use a context with timeout for the initial token request and potentially background refresh
	// // It's generally better to use the passed-in context 'ctx' for the client creation
	// // but ensure it doesn't cancel prematurely for background refreshes if needed.
	// // Using context.Background() for the token source might be safer for long-lived clients.
	// oauthCtx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{Timeout: 15 * time.Second}) // Use Background for token source
	// httpClient := conf.Client(oauthCtx)
	// s.httpClient = httpClient
	// --- END COMMENTED OUT OAuth2 CODE ---

	// Use a standard HTTP client since authentication is bypassed
	g.Log().Warning(ctx, "Wallet API client is configured WITHOUT OAuth2 authentication.")
	s.httpClient = &http.Client{Timeout: 15 * time.Second} // Use a default client with timeout

	// Configure the walletsclient
	walletsCfg := walletsclient.NewConfiguration()
	walletsCfg.HTTPClient = s.httpClient
	// Ensure the base URL ends with a slash if the paths in openapi.yaml don't start with one.
	// Based on the provided openapi.yaml, paths start with '/', so no trailing slash needed for baseURL.
	walletsCfg.Servers = walletsclient.ServerConfigurations{
		walletsclient.ServerConfiguration{
			URL:         s.baseURL, // Use the base URL from config
			Description: "Default Server",
		},
	}
	// Optional debug flag from config
	// walletsCfg.Debug = g.Cfg().MustGet(ctx, "walletsApi.debug", false).Bool()

	s.client = walletsclient.NewAPIClient(walletsCfg)

	g.Log().Infof(ctx, "Wallet API client initialized for base URL: %s", s.baseURL)
	return s.client, nil
}

// Note: Close method is usually not needed for http clients managed by oauth2 library.
// The underlying connections are typically managed by the http transport.
