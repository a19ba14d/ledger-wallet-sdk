package service

import (
	"context"

	"github.com/a19ba14d/ledger-wallet-sdk/internal/client"
	walletsclient "github.com/a19ba14d/ledger-wallet-sdk/internal/generated/v1"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// IWalletClient 定义了钱包客户端接口
type IWalletClient interface {
	GetClient(ctx context.Context) (*walletsclient.APIClient, error)
}

type sWalletClient struct {
	client *client.Client
}

var (
	insWalletClient IWalletClient
)

func init() {
	// 创建客户端配置
	cfg := client.Config{
		BaseURL: g.Cfg().MustGet(context.Background(), "wallet.baseURL", "http://localhost:8080").String(),
	}

	// 初始化客户端
	c, err := client.NewClient(cfg)
	if err != nil {
		g.Log().Errorf(context.Background(), "初始化钱包客户端失败: %v", err)
		// 仍然创建一个实例，但在使用时会返回错误
		insWalletClient = NewWalletClient(nil)
		return
	}

	// 初始化钱包客户端服务
	insWalletClient = NewWalletClient(c)
}

// NewWalletClient 创建一个新的钱包客户端服务实例
func NewWalletClient(client *client.Client) *sWalletClient {
	return &sWalletClient{
		client: client,
	}
}

// WalletClient 返回注册的钱包客户端服务实例
func WalletClient() IWalletClient {
	return insWalletClient
}

// GetClient 返回钱包API客户端
func (s *sWalletClient) GetClient(ctx context.Context) (*walletsclient.APIClient, error) {
	if s.client == nil {
		return nil, gerror.New("钱包客户端未正确初始化")
	}
	return s.client.GetGeneratedClient(), nil
}
