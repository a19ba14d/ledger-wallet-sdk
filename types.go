package ledgerwalletsdk

import "time"

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
	Asset  string `json:"asset"`  // 例如 "USD", "EUR"
	Amount int64  `json:"amount"` // 最小货币单位的数量
}

// Subject 代表一个交易的来源或目的地。
// 类型可以是 "ACCOUNT" 或 "WALLET"。
// Identifier 是账户 ID 或钱包 ID。
type Subject struct {
	Type       string `json:"type"` // "ACCOUNT" or "WALLET"
	Identifier string `json:"identifier"`
}

// Balance 代表特定资产的余额信息。
type Balance struct {
	Name   string    `json:"name"` // 余额名称 (通常是资产名称)
	Amount Monetary  `json:"amount"`
	// 可以根据需要添加其他余额相关字段，如 Expiry, Priority 等
}

// BalanceWithAssets 类似于 Balance，但可能包含更详细的资产信息（如果 API 支持）。
// 目前，我们将其定义为与 Balance 相同，如果需要扩展可以后续修改。
type BalanceWithAssets Balance

// Hold 代表一个冻结（预授权）操作。
type Hold struct {
	ID          string            `json:"id"`
	WalletID    string            `json:"walletID"`
	Description string            `json:"description"`
	Amount      Monetary          `json:"amount"` // 冻结的原始金额
	Remaining   Monetary          `json:"remaining"` // 冻结的剩余金额 (确认后减少)
	Metadata    map[string]string `json:"metadata"`
	Destination *Subject          `json:"destination,omitempty"` // 可选的目的地
	CreatedAt   time.Time         `json:"createdAt"`
	ExpiredAt   *time.Time        `json:"expiredAt,omitempty"` // 可选的过期时间
	IsVoid      bool              `json:"isVoid"` // 是否已作废
}

// ExpandedDebitHold 提供了更详细的冻结信息，通常包含关联的交易。
// 为了简化，我们先定义基础字段，如果需要交易信息，可以通过 GetTransaction 获取。
type ExpandedDebitHold Hold // 暂时使用 Hold，可后续扩展

// Transaction 代表一笔账本交易。
type Transaction struct {
	ID        int64             `json:"id"` // 注意：原始 API 可能是 int64
	Timestamp time.Time         `json:"timestamp"`
	Postings  []Posting         `json:"postings"`
	Reference *string           `json:"reference,omitempty"`
	Metadata  map[string]string `json:"metadata"`
}

// Posting 代表交易中的一个分录（借方或贷方）。
type Posting struct {
	Source      string   `json:"source"`
	Destination string   `json:"destination"`
	Amount      Monetary `json:"amount"`
}

// WalletSummary 提供了钱包的摘要信息。
type WalletSummary struct {
	AvailableFunds map[string]int64 `json:"availableFunds"` // 各资产可用余额
	ExpiredFunds   map[string]int64 `json:"expiredFunds"`   // 各资产已过期冻结金额
	HoldFunds      map[string]int64 `json:"holdFunds"`      // 各资产当前冻结金额
}

// ServerInfo 包含服务器版本等信息。
type ServerInfo struct {
	Version string `json:"version"`
}

// Cursor 用于分页查询的响应。
type Cursor[T any] struct {
	PageSize    int    `json:"pageSize"`
	HasMore     bool   `json:"hasMore"`
	Previous    *string `json:"previous,omitempty"`
	Next        *string `json:"next,omitempty"`
	Data        []T    `json:"data"`
}

// --- 辅助类型 ---

// Timestamp 只是 time.Time 的别名，如果需要可以移除。
type Timestamp = time.Time