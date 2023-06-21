package action

import (
	"context"
	"fmt"
	bankimpl "github.com/roseboy/bank-server/internal/bank-impl"
	"github.com/roseboy/go-ng/plugin"
)

// OpenAccountRequest Request
type OpenAccountRequest struct {
	BankType string
	Name     string
}

// OpenAccountResponse Response
type OpenAccountResponse struct {
	AccountNum string
}

// OpenAccount action
func OpenAccount(ctx context.Context, request, response any) error {
	meta := plugin.ExtractActionMeta(ctx)
	fmt.Printf("Login AppId: %d\n", meta.AppId)
	var req, resp = request.(*OpenAccountRequest), response.(*OpenAccountResponse)

	bankRet, err := bankimpl.CallBankOpenAccount(req.BankType, &bankimpl.OpenAccountArgs{})
	if err != nil {
		return err
	}

	resp.AccountNum = bankRet.AccountNum
	return nil
}
