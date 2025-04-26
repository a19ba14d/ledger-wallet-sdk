package service

import (
	"context"
	"net/http"
	"strings"

	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1" // 确认路径

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	// "github.com/shopspring/decimal" // Removed unused import
)

func (s *sWallet) CreateWallet(ctx context.Context, name string, metadata map[string]string) (*walletsclient.Wallet, error) {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	// Ensure metadata is not nil, as required by CreateWalletRequest
	if metadata == nil {
		metadata = make(map[string]string)
	}

	req := walletsclient.NewCreateWalletRequest(metadata, name)
	resp, httpResp, err := client.WalletsV1API.CreateWallet(ctx).CreateWalletRequest(*req).Execute()

	g.Log().Debug(ctx, "CreateWallet Request Payload: %+v", resp)

	// 错误处理
	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}

		// 临时处理：如果 HTTP 状态为 201 Created 但出现特定的 JSON 解析错误，则忽略
		if httpResp != nil && httpResp.StatusCode == http.StatusCreated && strings.Contains(err.Error(), "json: unknown field \"expiresAt\"") {
			g.Log().Warningf(ctx, "CreateWallet API 调用成功 (HTTP %d), 但响应解析时忽略了未知字段 'expiresAt': %v", httpResp.StatusCode, err)
			err = nil // 忽略此特定错误
		}

		// 如果错误仍然存在（不是被忽略的特定错误）
		if err != nil {
			g.Log().Errorf(ctx, "CreateWallet API 调用失败: %v, HTTP Status: %s", err, status)
			// Consider logging response body if needed: bodyBytes, _ := io.ReadAll(httpResp.Body)
			return nil, gerror.Wrapf(err, "创建钱包 API 调用失败 (HTTP: %s)", status)
		}
	}
	if resp == nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		return nil, gerror.Newf("创建钱包 API 返回空响应 (HTTP: %s)", status)
	}

	return &resp.Data, nil
}
