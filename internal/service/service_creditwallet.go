package service

import (
	"context"
	"net/http"
	"time"

	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1" // 确认路径

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	// "github.com/shopspring/decimal" // Removed unused import
)

func (s *sWallet) CreditWallet(ctx context.Context, walletID string, amount walletsclient.Monetary, sources []walletsclient.Subject, reference *string, metadata map[string]string, balance *string, timestamp *time.Time) error {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	// Validate amount: must be a positive integer
	if amount.Amount < 0 {
		return gerror.Newf("无效金额：金额必须是正整数 (传入值: %d)", amount.Amount)
	}

	req := walletsclient.NewCreditWalletRequest(amount)
	if sources != nil { // Ensure sources is not nil before setting
		req.SetSources(sources)
	}
	if reference != nil {
		req.SetReference(*reference)
	}
	if metadata != nil { // Ensure metadata is not nil before setting
		req.SetMetadata(metadata)
	}
	if balance != nil {
		req.SetBalance(*balance)
	}
	if timestamp != nil {
		req.SetTimestamp(*timestamp)
	}
	g.Log().Debug(ctx, "CreditWallet Request Payload (WalletID: %s): %+v", walletID, req)

	httpResp, err := client.WalletsV1API.CreditWallet(ctx, walletID).CreditWalletRequest(*req).Execute()

	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		g.Log().Errorf(ctx, "CreditWallet API 调用失败 (ID: %s): %v, HTTP Status: %s", walletID, err, status)
		return gerror.Wrapf(err, "钱包充值 API 调用失败 (ID: %s, HTTP: %s)", walletID, status)
	}
	if httpResp.StatusCode != http.StatusNoContent {
		g.Log().Warningf(ctx, "CreditWallet API 返回非预期成功状态码 (ID: %s): %s", walletID, httpResp.Status)
	}

	return nil
}
