package service

import (
	"context"

	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1" // 确认路径

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	// "github.com/shopspring/decimal" // Removed unused import
)

func (s *sWallet) GetWalletSummary(ctx context.Context, walletID string) (*walletsclient.WalletSummary, error) {
	// Get the client from the service struct
	apiClient, err := s.client.GetClient(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	resp, httpResp, err := apiClient.WalletsV1API.GetWalletSummary(ctx, walletID).Execute() // Use apiClient

	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		g.Log().Errorf(ctx, "GetWalletSummary API 调用失败 (ID: %s): %v, HTTP Status: %s", walletID, err, status)
		return nil, gerror.Wrapf(err, "获取钱包摘要 API 调用失败 (ID: %s, HTTP: %s)", walletID, status)
	}
	if resp == nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		return nil, gerror.Newf("获取钱包摘要 API 返回空响应 (ID: %s, HTTP: %s)", walletID, status)
	}

	return &resp.Data, nil
}
