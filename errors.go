package ledgerwalletsdk

import "errors"

var (
	// ErrWalletNotFound 示例错误
	ErrWalletNotFound = errors.New("wallet not found")
)

// Wrap 简易错误包装
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return errors.New(msg + ": " + err.Error())
}