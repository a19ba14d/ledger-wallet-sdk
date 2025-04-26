package service

import (
	"context"
	"net/http"

	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1" // 确认路径

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	// "github.com/shopspring/decimal" // Removed unused import
)

func (s *sWallet) UpdateWallet(ctx context.Context, walletID string, metadata map[string]string) error {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	// Ensure metadata is not nil, as required by UpdateWalletRequest
	if metadata == nil {
		metadata = make(map[string]string)
	}

	req := walletsclient.NewUpdateWalletRequest(metadata)
	httpResp, err := client.WalletsV1API.UpdateWallet(ctx, walletID).UpdateWalletRequest(*req).Execute()

	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		g.Log().Errorf(ctx, "UpdateWallet API 调用失败 (ID: %s): %v, HTTP Status: %s", walletID, err, status)
		return gerror.Wrapf(err, "更新钱包 API 调用失败 (ID: %s, HTTP: %s)", walletID, status)
	}
	// Check for non-204 status codes which might indicate success but have content, or other issues
	if httpResp.StatusCode != http.StatusNoContent {
		// Log or handle unexpected successful status codes if necessary
		g.Log().Warningf(ctx, "UpdateWallet API 返回非预期成功状态码 (ID: %s): %s", walletID, httpResp.Status)
	}

	return nil
}
