package ledgerwalletsdk

import (
	"context"
	"time"
)

// IWallet 定义了与钱包交互的核心操作接口。
// 所有方法都接收 context.Context 作为第一个参数。
type IWallet interface {
	// CreateWallet 创建一个新的钱包。
	CreateWallet(ctx context.Context, name string, metadata map[string]string) (*Wallet, error)

	// GetWallet 获取指定 ID 的钱包信息，包含余额摘要（如果可用）。
	// 注意：原始 API 返回 WalletWithBalances，这里简化为 Wallet，余额通过 ListBalances 获取。
	// 如果需要合并，可以调整返回类型或添加选项。
	GetWallet(ctx context.Context, walletID string) (*Wallet, error) // 返回我们定义的 Wallet

	// ListWallets 列出钱包，支持分页和过滤。
	ListWallets(ctx context.Context, params ListWalletsParams) (*Cursor[Wallet], error) // 返回 Cursor[Wallet]

	// UpdateWallet 更新指定钱包的元数据。
	UpdateWallet(ctx context.Context, walletID string, metadata map[string]string) error

	// CreditWallet 向指定钱包的特定余额（或默认余额）充值。
	CreditWallet(ctx context.Context, walletID string, amount Monetary, sources []Subject, reference *string, metadata map[string]string, balance *string, timestamp *time.Time) error // 使用我们定义的 Monetary 和 Subject

	// DebitWallet 从指定钱包扣款。
	// 如果 pending 为 true，则创建一个冻结 (Hold) 而不是立即扣款。
	// 返回创建的 Hold 信息（如果 pending=true），否则返回 nil。
	DebitWallet(ctx context.Context, walletID string, amount Monetary, pending *bool, metadata map[string]string, description *string, destination *Subject, balances []string, timestamp *time.Time) (*Hold, error) // 使用我们定义的 Monetary, Subject, Hold

	// GetBalance 获取钱包中特定余额的详细信息。
	GetBalance(ctx context.Context, walletID, balanceName string) (*BalanceWithAssets, error) // 使用我们定义的 BalanceWithAssets

	// ListBalances 列出钱包的所有余额。
	ListBalances(ctx context.Context, walletID string) (*Cursor[Balance], error) // 返回 Cursor[Balance]

	// GetHold 获取指定 ID 的冻结信息。
	GetHold(ctx context.Context, holdID string) (*ExpandedDebitHold, error) // 使用我们定义的 ExpandedDebitHold

	// ListHolds 列出冻结，支持分页和按钱包 ID 过滤。
	ListHolds(ctx context.Context, params ListHoldsParams) (*Cursor[Hold], error) // 返回 Cursor[Hold]

	// ConfirmHold 确认一个冻结，将其转换为实际扣款。
	// amount 可选，指定确认的金额（小于等于冻结金额）。
	// final 可选，指示这是否是最后一次确认。
	ConfirmHold(ctx context.Context, holdID string, amount *int64, final *bool) error // amount 改为 int64

	// VoidHold 作废一个冻结，释放预留的资金。
	VoidHold(ctx context.Context, holdID string) error

	// GetServerInfo 获取服务器信息。
	GetServerInfo(ctx context.Context) (*ServerInfo, error) // 使用我们定义的 ServerInfo

	// GetTransactions 列出交易记录，支持分页和按钱包 ID 过滤。
	GetTransactions(ctx context.Context, params GetTransactionsParams) (*Cursor[Transaction], error) // 返回 Cursor[Transaction]

	// GetWalletSummary 获取钱包的资金摘要信息。
	GetWalletSummary(ctx context.Context, walletID string) (*WalletSummary, error) // 使用我们定义的 WalletSummary

	// CreateBalance 在钱包中创建一个新的余额账户（例如用于特定资产）。
	CreateBalance(ctx context.Context, walletID string, balance Balance) (*Balance, error) // 使用我们定义的 Balance
}