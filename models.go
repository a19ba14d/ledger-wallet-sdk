package ledgerwalletsdk

// ListWalletsParams 查询钱包列表参数
type ListWalletsParams struct {
	Name     *string
	Metadata map[string]string
	PageSize *int32
	Cursor   *string
	Expand   []string
}

// ListHoldsParams 查询冻结列表参数
type ListHoldsParams struct {
	WalletID *string
	Metadata map[string]string
	PageSize *int32
	Cursor   *string
}

// GetTransactionsParams 查询交易列表参数
type GetTransactionsParams struct {
	PageSize *int32
	WalletID *string
	Cursor   *string
}

// Timestamp alias moved to types.go