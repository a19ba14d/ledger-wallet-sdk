package ledgerwalletsdk

import (
	"time"

	v1 "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1"
)

// Wallet 代表一个钱包资源。
type Wallet struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Metadata  map[string]string `json:"metadata"`
	CreatedAt time.Time         `json:"createdAt"`
	// 注意：省略了 Ledger 字段，如果需要暴露，可以后续添加。
	// 注意：Balances 将通过专门的方法获取，不直接包含在此结构中以简化。
}

// Monetary 代表一个货币金额，包含资产类型和数量。
// Amount 使用 int64 表示最小货币单位（例如美分的数量）。
type Monetary struct {
	*v1.Monetary
}

// Subject 代表一个交易的来源或目的地。
// 类型可以是 "ACCOUNT" 或 "WALLET"。
// Identifier 是账户 ID 或钱包 ID。
type Subject struct {
	*v1.Subject
}

// Balance 代表特定资产的余额信息。
type Balance struct {
	*v1.Balance
}

// BalanceWithAssets 类似于 Balance，但可能包含更详细的资产信息（如果 API 支持）。
// 目前，我们将其定义为与 Balance 相同，如果需要扩展可以后续修改。
type BalanceWithAssets Balance

// Hold 代表一个冻结（预授权）操作。
type Hold struct {
	*v1.Hold
}

// ExpandedDebitHold 提供了更详细的冻结信息，通常包含关联的交易。
// 为了简化，我们先定义基础字段，如果需要交易信息，可以通过 GetTransaction 获取。
type ExpandedDebitHold Hold // 暂时使用 Hold，可后续扩展

// Transaction 代表一笔账本交易。
type Transaction struct {
	*v1.Transaction
}

// Posting 代表交易中的一个分录（借方或贷方）。
type Posting struct {
	*v1.Posting
}

// WalletSummary 提供了钱包的摘要信息。
type WalletSummary struct {
	*v1.WalletSummary
}

// ServerInfo 包含服务器版本等信息。
type ServerInfo struct {
	*v1.ServerInfo
}

// Cursor 用于分页查询的响应。
type Cursor struct {
	*v1.Cursor
}

// --- 辅助类型 ---

// Timestamp 只是 time.Time 的别名，如果需要可以移除。
type Timestamp = time.Time
