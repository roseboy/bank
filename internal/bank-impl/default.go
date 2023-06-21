package bankimpl

import "fmt"

// BankImpl Bank
type BankImpl interface {
	OpenAccount(*OpenAccountArgs) (*OpenAccountRet, error)
}

// OpenAccountArgs args
type OpenAccountArgs struct {
	Name  string
	IdNum string
}

// OpenAccountRet ret
type OpenAccountRet struct {
	AccountNum string
}

var bankImplMap = make(map[string]BankImpl)

// RegisterBank Register
func RegisterBank(bankType string, bank BankImpl) {
	bankImplMap[bankType] = bank
}

// CallBankOpenAccount call
func CallBankOpenAccount(bankType string, args *OpenAccountArgs) (*OpenAccountRet, error) {
	bankImpl, ok := bankImplMap[bankType]
	if !ok {
		return nil, fmt.Errorf("not found bank:%s", bankType)
	}
	return bankImpl.OpenAccount(args)
}
