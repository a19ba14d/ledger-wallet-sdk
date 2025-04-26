package ledgerwalletsdk

// Timestamp alias moved to types.go

// ListWalletsParams holds parameters for listing wallets.
type ListWalletsParams struct {
	Name     *string
	Metadata map[string]string
	PageSize *int32
	Cursor   *string
	Expand   []string
}

// ListHoldsParams holds parameters for listing holds.
type ListHoldsParams struct {
	WalletID *string
	Metadata map[string]string
	PageSize *int32
	Cursor   *string
}

// GetTransactionsParams holds parameters for listing transactions.
type GetTransactionsParams struct {
	PageSize *int32
	WalletID *string
	Cursor   *string
}
