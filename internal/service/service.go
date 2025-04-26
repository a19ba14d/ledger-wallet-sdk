package service

import (
	"bytes"
	"context"
	"encoding/json"
"fmt"
	"strings"
	"io"
	"net/http"
	"time"

	walletsclient "telegram-bot-api/internal/client/wallets/v1" // 确认路径

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	// "github.com/shopspring/decimal" // Removed unused import
)

// ListWalletsParams holds parameters for listing wallets.
type ListWalletsParams struct {
	Name     *string
	Metadata map[string]string
	PageSize *int32
	Cursor   *string
	Expand   []string
}

// ListHoldsParams holds parameters for listing holds.
type ListHoldsParams struct {
	WalletID *string
	Metadata map[string]string
	PageSize *int32
	Cursor   *string
}

// GetTransactionsParams holds parameters for listing transactions.
type GetTransactionsParams struct {
	PageSize *int32
	WalletID *string
	Cursor   *string
}

// IWallet defines the interface for wallet service operations.
type IWallet interface {
	CreateWallet(ctx context.Context, name string, metadata map[string]string) (*walletsclient.Wallet, error)
	GetWallet(ctx context.Context, walletID string) (*walletsclient.WalletWithBalances, error)
	ListWallets(ctx context.Context, params ListWalletsParams) (*walletsclient.ListWalletsResponseCursor, error)
	UpdateWallet(ctx context.Context, walletID string, metadata map[string]string) error
	CreditWallet(ctx context.Context, walletID string, amount walletsclient.Monetary, sources []walletsclient.Subject, reference *string, metadata map[string]string, balance *string, timestamp *time.Time) error
	DebitWallet(ctx context.Context, walletID string, amount walletsclient.Monetary, pending *bool, metadata map[string]string, description *string, destination *walletsclient.Subject, balances []string, timestamp *time.Time) (*walletsclient.Hold, error) // Returns Hold if pending=true
	GetBalance(ctx context.Context, walletID, balanceName string) (*walletsclient.BalanceWithAssets, error)
	ListBalances(ctx context.Context, walletID string) (*walletsclient.ListBalancesResponseCursor, error)
	GetHold(ctx context.Context, holdID string) (*walletsclient.ExpandedDebitHold, error)
	ListHolds(ctx context.Context, params ListHoldsParams) (*walletsclient.GetHoldsResponseCursor, error)
	ConfirmHold(ctx context.Context, holdID string, amount *int32, final *bool) error
	VoidHold(ctx context.Context, holdID string) error
	// Add other methods as needed based on walletsclient capabilities and business logic
	GetServerInfo(ctx context.Context) (*walletsclient.ServerInfo, error)
	GetTransactions(ctx context.Context, params GetTransactionsParams) (*walletsclient.GetTransactionsResponseCursor, error)
	GetWalletSummary(ctx context.Context, walletID string) (*walletsclient.WalletSummary, error)
	CreateBalance(ctx context.Context, walletID string, balance walletsclient.Balance) (*walletsclient.Balance, error)
	// Removed GetUserTokenBalance, ConvertBalanceToRaw, ConvertRawToBalance, FormatUserBalance, FormatUserBalanceWithSymbol, GetFormattedUserTokenBalance, GetFormattedUserTokenBalanceWithSymbol
}

type sWallet struct {
	// No direct dependency needed here if we always get client via WalletClient()
	// walletClient IWalletClient
}

var (
	insWallet IWallet
)

func init() {
	insWallet = NewWallet()
}

// NewWallet creates and returns a new instance of the wallet service.
func NewWallet() *sWallet {
	return &sWallet{}
}

// Wallet returns the registered wallet service instance.
func Wallet() IWallet {
	// No need for panic check if using init() registration
	return insWallet
}

// --- Implementation of IWallet ---

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

func (s *sWallet) GetWallet(ctx context.Context, walletID string) (*walletsclient.WalletWithBalances, error) {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	resp, httpResp, err := client.WalletsV1API.GetWallet(ctx, walletID).Execute()

	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		g.Log().Errorf(ctx, "GetWallet API 调用失败 (ID: %s): %v, HTTP Status: %s", walletID, err, status)
		return nil, gerror.Wrapf(err, "获取钱包信息 API 调用失败 (ID: %s, HTTP: %s)", walletID, status)
	}
	if resp == nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		return nil, gerror.Newf("获取钱包信息 API 返回空响应 (ID: %s, HTTP: %s)", walletID, status)
	}

	return &resp.Data, nil
}

func (s *sWallet) ListWallets(ctx context.Context, params ListWalletsParams) (*walletsclient.ListWalletsResponseCursor, error) {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	apiReq := client.WalletsV1API.ListWallets(ctx)
	if params.Name != nil {
		apiReq = apiReq.Name(*params.Name)
	}
	if params.Metadata != nil {
		apiReq = apiReq.Metadata(params.Metadata)
	}
	if params.PageSize != nil {
		apiReq = apiReq.PageSize(int64(*params.PageSize))
	}
	if params.Cursor != nil {
		apiReq = apiReq.Cursor(*params.Cursor)
	}
	if params.Expand != nil {
		apiReq = apiReq.Expand(params.Expand)
	}

	resp, httpResp, err := apiReq.Execute()

	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		g.Log().Errorf(ctx, "ListWallets API 调用失败: %v, HTTP Status: %s", err, status)
		return nil, gerror.Wrapf(err, "列出钱包 API 调用失败 (HTTP: %s)", status)
	}
	if resp == nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		return nil, gerror.Newf("列出钱包 API 返回空响应 (HTTP: %s)", status)
	}

	return &resp.Cursor, nil
}

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

// DebitWallet debits a wallet, potentially creating a hold if pending is true.
func (s *sWallet) DebitWallet(ctx context.Context, walletID string, amount walletsclient.Monetary, pending *bool, metadata map[string]string, description *string, destination *walletsclient.Subject, balances []string, timestamp *time.Time) (*walletsclient.Hold, error) {
	client, err := WalletClient().GetClient(ctx)
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

	resp, httpResp, err := client.WalletsV1API.DebitWallet(ctx, walletID).DebitWalletRequest(*req).Execute()

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

func (s *sWallet) GetBalance(ctx context.Context, walletID, balanceName string) (*walletsclient.BalanceWithAssets, error) {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	resp, httpResp, err := client.WalletsV1API.GetBalance(ctx, walletID, balanceName).Execute()

	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		g.Log().Errorf(ctx, "GetBalance API 调用失败 (WalletID: %s, Balance: %s): %v, HTTP Status: %s", walletID, balanceName, err, status)
		return nil, gerror.Wrapf(err, "获取余额详情 API 调用失败 (WalletID: %s, Balance: %s, HTTP: %s)", walletID, balanceName, status)
	}
	if resp == nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		return nil, gerror.Newf("获取余额详情 API 返回空响应 (WalletID: %s, Balance: %s, HTTP: %s)", walletID, balanceName, status)
	}

	return &resp.Data, nil
}

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

func (s *sWallet) GetHold(ctx context.Context, holdID string) (*walletsclient.ExpandedDebitHold, error) {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	resp, httpResp, err := client.WalletsV1API.GetHold(ctx, holdID).Execute()

	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		g.Log().Errorf(ctx, "GetHold API 调用失败 (HoldID: %s): %v, HTTP Status: %s", holdID, err, status)
		return nil, gerror.Wrapf(err, "获取冻结信息 API 调用失败 (HoldID: %s, HTTP: %s)", holdID, status)
	}
	if resp == nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		return nil, gerror.Newf("获取冻结信息 API 返回空响应 (HoldID: %s, HTTP: %s)", holdID, status)
	}

	return &resp.Data, nil
}

func (s *sWallet) ListHolds(ctx context.Context, params ListHoldsParams) (*walletsclient.GetHoldsResponseCursor, error) {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	apiReq := client.WalletsV1API.GetHolds(ctx)
	if params.WalletID != nil {
		apiReq = apiReq.WalletID(*params.WalletID)
	}
	if params.Metadata != nil {
		apiReq = apiReq.Metadata(params.Metadata)
	}
	if params.PageSize != nil {
		apiReq = apiReq.PageSize(int64(*params.PageSize))
	}
	if params.Cursor != nil {
		apiReq = apiReq.Cursor(*params.Cursor)
	}

	resp, httpResp, err := apiReq.Execute()

	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		g.Log().Errorf(ctx, "ListHolds API 调用失败: %v, HTTP Status: %s", err, status)
		return nil, gerror.Wrapf(err, "列出冻结 API 调用失败 (HTTP: %s)", status)
	}
	if resp == nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		return nil, gerror.Newf("列出冻结 API 返回空响应 (HTTP: %s)", status)
	}

	return &resp.Cursor, nil
}

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

func (s *sWallet) GetTransactions(ctx context.Context, params GetTransactionsParams) (*walletsclient.GetTransactionsResponseCursor, error) {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	apiReq := client.WalletsV1API.GetTransactions(ctx)
	if params.PageSize != nil {
		apiReq = apiReq.PageSize(int64(*params.PageSize))
	}
	if params.WalletID != nil {
		apiReq = apiReq.WalletID(*params.WalletID)
	}
	if params.Cursor != nil {
		apiReq = apiReq.Cursor(*params.Cursor)
	}

	resp, httpResp, err := apiReq.Execute()

	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		g.Log().Errorf(ctx, "GetTransactions API 调用失败: %v, HTTP Status: %s", err, status)
		return nil, gerror.Wrapf(err, "获取交易记录 API 调用失败 (HTTP: %s)", status)
	}
	if resp == nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		return nil, gerror.Newf("获取交易记录 API 返回空响应 (HTTP: %s)", status)
	}

	return &resp.Cursor, nil
}

func (s *sWallet) GetWalletSummary(ctx context.Context, walletID string) (*walletsclient.WalletSummary, error) {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	resp, httpResp, err := client.WalletsV1API.GetWalletSummary(ctx, walletID).Execute()

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

func (s *sWallet) CreateBalance(ctx context.Context, walletID string, balance walletsclient.Balance) (*walletsclient.Balance, error) {
	client, err := WalletClient().GetClient(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "获取 Wallet API 客户端失败")
	}

	req := balance // Use the provided balance object directly as the request body
	resp, httpResp, err := client.WalletsV1API.CreateBalance(ctx, walletID).Body(req).Execute()

	if err != nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		g.Log().Errorf(ctx, "CreateBalance API 调用失败 (WalletID: %s): %v, HTTP Status: %s", walletID, err, status)
		return nil, gerror.Wrapf(err, "创建余额 API 调用失败 (WalletID: %s, HTTP: %s)", walletID, status)
	}
	if resp == nil {
		status := "N/A"
		if httpResp != nil {
			status = httpResp.Status
		}
		return nil, gerror.Newf("创建余额 API 返回空响应 (WalletID: %s, HTTP: %s)", walletID, status)
	}

	return &resp.Data, nil
}

// GetUserTokenBalance returns a user's token balance for a specific symbol as a decimal.Decimal.
// Removed GetUserTokenBalance, ConvertBalanceToRaw, ConvertRawToBalance, FormatUserBalance, FormatUserBalanceWithSymbol, GetFormattedUserTokenBalance, GetFormattedUserTokenBalanceWithSymbol implementations
