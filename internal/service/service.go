package service

import (
	"context"
	"time"

	// Removed import of root package ledgerwalletsdk
	internalClient "github.com/a19ba14d/ledger-wallet-sdk/internal/client"      // Import internal client package
	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1" // Keep generated client
	sdkTypes "github.com/a19ba14d/ledger-wallet-sdk/pkg/types"                   // Import shared types package
	// "github.com/shopspring/decimal" // Removed unused import
)

// IWallet defines the interface for wallet service operations.
type IWallet interface {
	CreateWallet(ctx context.Context, name string, metadata map[string]string) (*walletsclient.Wallet, error)
	GetWallet(ctx context.Context, walletID string) (*walletsclient.WalletWithBalances, error)
	ListWallets(ctx context.Context, params sdkTypes.ListWalletsParams) (*walletsclient.ListWalletsResponseCursor, error) // Use sdkTypes
	UpdateWallet(ctx context.Context, walletID string, metadata map[string]string) error
	CreditWallet(ctx context.Context, walletID string, amount walletsclient.Monetary, sources []walletsclient.Subject, reference *string, metadata map[string]string, balance *string, timestamp *time.Time) error
	DebitWallet(ctx context.Context, walletID string, amount walletsclient.Monetary, pending *bool, metadata map[string]string, description *string, destination *walletsclient.Subject, balances []string, timestamp *time.Time) (*walletsclient.Hold, error) // Returns Hold if pending=true
	GetBalance(ctx context.Context, walletID, balanceName string) (*walletsclient.BalanceWithAssets, error)
	ListBalances(ctx context.Context, walletID string) (*walletsclient.ListBalancesResponseCursor, error)
	GetHold(ctx context.Context, holdID string) (*walletsclient.ExpandedDebitHold, error)
	ListHolds(ctx context.Context, params sdkTypes.ListHoldsParams) (*walletsclient.GetHoldsResponseCursor, error) // Use sdkTypes
	ConfirmHold(ctx context.Context, holdID string, amount *int32, final *bool) error
	VoidHold(ctx context.Context, holdID string) error
	// Add other methods as needed based on walletsclient capabilities and business logic
	GetServerInfo(ctx context.Context) (*walletsclient.ServerInfo, error)
	GetTransactions(ctx context.Context, params sdkTypes.GetTransactionsParams) (*walletsclient.GetTransactionsResponseCursor, error) // Use sdkTypes
	GetWalletSummary(ctx context.Context, walletID string) (*walletsclient.WalletSummary, error)
	CreateBalance(ctx context.Context, walletID string, balance walletsclient.Balance) (*walletsclient.Balance, error)
	// Removed GetUserTokenBalance, ConvertBalanceToRaw, ConvertRawToBalance, FormatUserBalance, FormatUserBalanceWithSymbol, GetFormattedUserTokenBalance, GetFormattedUserTokenBalanceWithSymbol
}

// sWallet implements the IWallet interface and holds the client dependency.
type sWallet struct {
	client internalClient.IWalletClient // Holds the injected client instance
}

// NewWallet creates and returns a new instance of the wallet service,
// accepting the wallet client as a dependency.
func NewWallet(client internalClient.IWalletClient) *sWallet {
	return &sWallet{
		client: client,
	}
}

// Removed init(), insWallet, and Wallet() function as per the plan.
// The service instance will now be created and managed by the SDK's New function.
