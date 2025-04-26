package service

import (
	"context"
	"net/http"

	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1" // 确认路径

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	// "github.com/shopspring/decimal" // Removed unused import
)

func (s *sWallet) ConfirmHold(ctx context.Context, holdID string, amount *int32, final *bool) error {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	req := walletsclient.NewConfirmHoldRequest()
	if amount != nil {
		req.SetAmount(int64(*amount))
	}
	if final != nil {
		req.SetFinal(*final)
	}

	httpResp, err := client.WalletsV1API.ConfirmHold(ctx, holdID).ConfirmHoldRequest(*req).Execute()

	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		g.Log().Errorf(ctx, "ConfirmHold API 调用失败 (HoldID: %s): %v, HTTP Status: %s", holdID, err, status)
		return gerror.Wrapf(err, "确认冻结 API 调用失败 (HoldID: %s, HTTP: %s)", holdID, status)
	}
	if httpResp.StatusCode != http.StatusNoContent {
		g.Log().Warningf(ctx, "ConfirmHold API 返回非预期成功状态码 (HoldID: %s): %s", holdID, httpResp.Status)
	}

	return nil
}
