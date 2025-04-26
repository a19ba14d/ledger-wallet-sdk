package service

import (
	"context"
	"time"

	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1" // 确认路径
	models "github.com/a19ba14d/ledger-wallet-sdk/models"                       // 确认路径
	// "github.com/shopspring/decimal" // Removed unused import
)

// IWallet defines the interface for wallet service operations.
type IWallet interface {
	CreateWallet(ctx context.Context, name string, metadata map[string]string) (*walletsclient.Wallet, error)
	GetWallet(ctx context.Context, walletID string) (*walletsclient.WalletWithBalances, error)
	ListWallets(ctx context.Context, params models.ListWalletsParams) (*walletsclient.ListWalletsResponseCursor, error)
	UpdateWallet(ctx context.Context, walletID string, metadata map[string]string) error
	CreditWallet(ctx context.Context, walletID string, amount walletsclient.Monetary, sources []walletsclient.Subject, reference *string, metadata map[string]string, balance *string, timestamp *time.Time) error
	DebitWallet(ctx context.Context, walletID string, amount walletsclient.Monetary, pending *bool, metadata map[string]string, description *string, destination *walletsclient.Subject, balances []string, timestamp *time.Time) (*walletsclient.Hold, error) // Returns Hold if pending=true
	GetBalance(ctx context.Context, walletID, balanceName string) (*walletsclient.BalanceWithAssets, error)
	ListBalances(ctx context.Context, walletID string) (*walletsclient.ListBalancesResponseCursor, error)
	GetHold(ctx context.Context, holdID string) (*walletsclient.ExpandedDebitHold, error)
	ListHolds(ctx context.Context, params models.ListHoldsParams) (*walletsclient.GetHoldsResponseCursor, error)
	ConfirmHold(ctx context.Context, holdID string, amount *int32, final *bool) error
	VoidHold(ctx context.Context, holdID string) error
	// Add other methods as needed based on walletsclient capabilities and business logic
	GetServerInfo(ctx context.Context) (*walletsclient.ServerInfo, error)
	GetTransactions(ctx context.Context, params models.GetTransactionsParams) (*walletsclient.GetTransactionsResponseCursor, error)
	GetWalletSummary(ctx context.Context, walletID string) (*walletsclient.WalletSummary, error)
	CreateBalance(ctx context.Context, walletID string, balance walletsclient.Balance) (*walletsclient.Balance, error)
	// Removed GetUserTokenBalance, ConvertBalanceToRaw, ConvertRawToBalance, FormatUserBalance, FormatUserBalanceWithSymbol, GetFormattedUserTokenBalance, GetFormattedUserTokenBalanceWithSymbol
}

type sWallet struct {
}

var (
	insWallet IWallet
)

func init() {
	insWallet = NewWallet()
}

// NewWallet creates and returns a new instance of the wallet service.
func NewWallet() *sWallet {
	return &sWallet{}
}

// Wallet returns the registered wallet service instance.
func Wallet() IWallet {
	// No need for panic check if using init() registration
	return insWallet
}
