package service

import (
	"context"
	"net/http"

	// 确认路径

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	// "github.com/shopspring/decimal" // Removed unused import
)

func (s *sWallet) VoidHold(ctx context.Context, holdID string) error {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	httpResp, err := client.WalletsV1API.VoidHold(ctx, holdID).Execute()

	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		g.Log().Errorf(ctx, "VoidHold API 调用失败 (HoldID: %s): %v, HTTP Status: %s", holdID, err, status)
		return gerror.Wrapf(err, "取消冻结 API 调用失败 (HoldID: %s, HTTP: %s)", holdID, status)
	}
	if httpResp.StatusCode != http.StatusNoContent {
		g.Log().Warningf(ctx, "VoidHold API 返回非预期成功状态码 (HoldID: %s): %s", holdID, httpResp.Status)
	}

	return nil
}
