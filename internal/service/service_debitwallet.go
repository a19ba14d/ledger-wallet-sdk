package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1" // 确认路径

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	// "github.com/shopspring/decimal" // Removed unused import
)

// DebitWallet debits a wallet, potentially creating a hold if pending is true.
func (s *sWallet) DebitWallet(ctx context.Context, walletID string, amount walletsclient.Monetary, pending *bool, metadata map[string]string, description *string, destination *walletsclient.Subject, balances []string, timestamp *time.Time) (*walletsclient.Hold, error) {
	// Get the client from the service struct
	apiClient, err := s.client.GetClient(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	// Validate amount: must be a positive integer
	if amount.Amount < 0 {
		return nil, gerror.Newf("无效金额：金额必须是正整数 (传入值: %d)", amount.Amount)
	}

	// Metadata is required by the client, ensure it's not nil
	if metadata == nil {
		metadata = make(map[string]string) // Initialize if nil
	}

	req := walletsclient.NewDebitWalletRequest(amount, metadata)
	// if pending != nil {
	// 	req.SetPending(*pending)
	// }
	if description != nil {
		req.SetDescription(*description)
	}
	if destination != nil {
		req.SetDestination(*destination)
	}
	if balances != nil { // Ensure balances is not nil before setting
		req.SetBalances(balances)
	}
	if timestamp != nil {
		req.SetTimestamp(*timestamp)
	}

	// Log the request payload before sending
	reqJSON, jsonErr := json.Marshal(req)
	if jsonErr != nil {
		g.Log().Warningf(ctx, "Failed to marshal DebitWalletRequest to JSON for logging: %v", jsonErr)
	} else {
		g.Log().Debugf(ctx, "DebitWallet Request Payload (WalletID: %s): %s", walletID, string(reqJSON))
	}

	resp, httpResp, err := apiClient.WalletsV1API.DebitWallet(ctx, walletID).DebitWalletRequest(*req).Execute() // Use apiClient

	g.Log().Debug(ctx, "DebitWallet API 调用完成 (ID: %s)", resp)

	if err != nil {
		status := "N/A"
		respBody := "N/A"
		respHeaders := "N/A"

		if httpResp != nil {
			status = httpResp.Status
			// Read response body
			bodyBytes, readErr := io.ReadAll(httpResp.Body)
			if readErr != nil {
				respBody = fmt.Sprintf("Failed to read response body: %v", readErr)
			} else {
				respBody = string(bodyBytes)
				// Restore the body so it can be read again if needed elsewhere (though unlikely in this error path)
				httpResp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
			// Format headers
			headerBytes, jsonErr := json.Marshal(httpResp.Header)
			if jsonErr != nil {
				respHeaders = fmt.Sprintf("Failed to marshal headers: %v", jsonErr)
			} else {
				respHeaders = string(headerBytes)
			}
		}

		g.Log().Errorf(ctx, "DebitWallet API 调用失败 (ID: %s): %v\nHTTP Status: %s\nResponse Headers: %s\nResponse Body: %s",
			walletID, err, status, respHeaders, respBody)

		// Wrap the original error, potentially adding more context if needed
		return nil, gerror.Wrapf(err, "钱包扣款/冻结 API 调用失败 (ID: %s, HTTP: %s)", walletID, status)
	}

	// If pending=true, API returns 201 with Hold details
	// if pending != nil && *pending {
	// 	if resp == nil {
	// 		status := "N/A"
	// 		if httpResp != nil {
	// 			status = httpResp.Status
	// 		}
	// 		return nil, gerror.Newf("钱包冻结 API 未返回冻结信息 (ID: %s, HTTP: %s)", walletID, status)
	// 	}
	// 	return &resp.Data, nil
	// }

	// If pending=false or nil, API returns 204 on success
	if httpResp.StatusCode != http.StatusNoContent {
		g.Log().Warningf(ctx, "DebitWallet API (direct debit) 返回非预期成功状态码 (ID: %s): %s", walletID, httpResp.Status)
	}
	return nil, nil // No Hold object returned for direct debit
}
