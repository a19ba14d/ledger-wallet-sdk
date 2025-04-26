package service

import (
	"context"

	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1" // 确认路径

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	// "github.com/shopspring/decimal" // Removed unused import
)

func (s *sWallet) ListBalances(ctx context.Context, walletID string) (*walletsclient.ListBalancesResponseCursor, error) {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	resp, httpResp, err := client.WalletsV1API.ListBalances(ctx, walletID).Execute()

	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		g.Log().Errorf(ctx, "ListBalances API 调用失败 (WalletID: %s): %v, HTTP Status: %s", walletID, err, status)
		return nil, gerror.Wrapf(err, "列出余额 API 调用失败 (WalletID: %s, HTTP: %s)", walletID, status)
	}
	if resp == nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		return nil, gerror.Newf("列出余额 API 返回空响应 (WalletID: %s, HTTP: %s)", walletID, status)
	}

	return &resp.Cursor, nil
}
