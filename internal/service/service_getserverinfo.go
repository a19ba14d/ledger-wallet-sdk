package service

import (
	"context"

	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1" // 确认路径

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	// "github.com/shopspring/decimal" // Removed unused import
)

func (s *sWallet) GetServerInfo(ctx context.Context) (*walletsclient.ServerInfo, error) {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	resp, httpResp, err := client.WalletsV1API.GetServerInfo(ctx).Execute()

	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		g.Log().Errorf(ctx, "GetServerInfo API 调用失败: %v, HTTP Status: %s", err, status)
		return nil, gerror.Wrapf(err, "获取服务器信息 API 调用失败 (HTTP: %s)", status)
	}
	if resp == nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		return nil, gerror.Newf("获取服务器信息 API 返回空响应 (HTTP: %s)", status)
	}

	return resp, nil
}
