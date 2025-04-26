package ledgerwalletsdk

import (
	"sync"

	internalClient "github.com/a19ba14d/ledger-wallet-sdk/internal/client"
	internalService "github.com/a19ba14d/ledger-wallet-sdk/internal/service"
	"github.com/a19ba14d/ledger-wallet-sdk/pkg/sdkconfig" // Import the new config package
	// "github.com/a19ba14d/ledger-wallet-sdk/pkg/log" // Logger can be configured via options
)

var (
	sdkInstance IWallet   // Singleton instance of the wallet service
	sdkOnce     sync.Once // Ensures initialization happens only once
	initErr     error     // Stores any error during initialization
)

// IWallet is an alias for the internal service interface.
// This makes it available at the SDK root level.
type IWallet = internalService.IWallet

// New 使用函数式选项初始化并返回一个单例的 IWallet 服务实例。
// 第一次调用时会进行初始化，后续调用将返回同一个实例。
// 必须提供 WithBaseURL 选项。
func New(opts ...sdkconfig.Option) (IWallet, error) { // Updated Option type
	sdkOnce.Do(func() {
		// 1. Create default config using the function from sdkconfig
		config := sdkconfig.NewDefaultConfig() // Updated function call

		// 2. Apply provided options
		for _, opt := range opts {
			opt(config)
		}

		// 3. Validate configuration
		if err := config.Validate(); err != nil {
			// config.GetLogger().Errorf(context.Background(), "SDK configuration validation failed: %v", err) // Requires context and logger access method
			initErr = err // Store error
			return        // Exit Do func
		}

		// 4. Create Wallet Client, passing the config from sdkconfig
		client, err := internalClient.NewWalletClient(config) // Pass sdkconfig.Config
		if err != nil {
			// config.GetLogger().Errorf(context.Background(), "Failed to initialize wallet client: %v", err)
			initErr = err // Store error
			return        // Exit Do func
		}

		// 5. Create Wallet Service (injecting client)
		// Note: We are casting the concrete *sWallet to the IWallet interface here.
		sdkInstance = internalService.NewWallet(client)

		// Optional: Log successful initialization
		// config.GetLogger().Infof(context.Background(), "Ledger Wallet SDK initialized successfully.")
	})

	// Return the initialized instance (or nil if init failed) and the error status
	return sdkInstance, initErr
}
