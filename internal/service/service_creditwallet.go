package service

import (
	"context"
	"net/http"
	"time"

	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1" // 确认路径
	sdkTypes "github.com/a19ba14d/ledger-wallet-sdk/pkg/types"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	// "github.com/shopspring/decimal" // Removed unused import
)

func (s *sWallet) CreditWallet(ctx context.Context, walletID string, amount sdkTypes.Monetary, sources []walletsclient.Subject, reference *string, metadata map[string]string, balance *string, timestamp *time.Time) error {
	// Get the client from the service struct
	apiClient, err := s.client.GetClient(ctx)
	if err != nil {
		return gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	// Validate amount: must be a positive integer
	if amount.Amount < 0 {
		return gerror.Newf("无效金额：金额必须是正整数 (传入值: %d)", amount.Amount)
	}

	// 将 sdkTypes.Monetary 转换为 walletsclient.Monetary
	walletAmount := walletsclient.Monetary{
		Asset:  amount.Asset,
		Amount: amount.Amount,
	}
	
	req := walletsclient.NewCreditWalletRequest(walletAmount)
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

	httpResp, err := apiClient.WalletsV1API.CreditWallet(ctx, walletID).CreditWalletRequest(*req).Execute() // Use apiClient

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
