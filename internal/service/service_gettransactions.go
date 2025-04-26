package service

import (
	"context"

	ledgerwalletsdk "github.com/a19ba14d/ledger-wallet-sdk"
	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1" // 确认路径

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	// "github.com/shopspring/decimal" // Removed unused import
)

func (s *sWallet) GetTransactions(ctx context.Context, params ledgerwalletsdk.GetTransactionsParams) (*walletsclient.GetTransactionsResponseCursor, error) {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	apiReq := client.WalletsV1API.GetTransactions(ctx)
	if params.PageSize != nil {
		apiReq = apiReq.PageSize(int64(*params.PageSize))
	}
	if params.WalletID != nil {
		apiReq = apiReq.WalletID(*params.WalletID)
	}
	if params.Cursor != nil {
		apiReq = apiReq.Cursor(*params.Cursor)
	}

	resp, httpResp, err := apiReq.Execute()

	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		g.Log().Errorf(ctx, "GetTransactions API 调用失败: %v, HTTP Status: %s", err, status)
		return nil, gerror.Wrapf(err, "获取交易记录 API 调用失败 (HTTP: %s)", status)
	}
	if resp == nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		return nil, gerror.Newf("获取交易记录 API 返回空响应 (HTTP: %s)", status)
	}

	return &resp.Cursor, nil
}
