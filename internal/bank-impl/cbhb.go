package bankimpl

// CBHB 渤海银行
type CBHB struct {
}

func init() {
	RegisterBank("CBHB", new(CBHB))
}

// OpenAccount open
func (C *CBHB) OpenAccount(args *OpenAccountArgs) (*OpenAccountRet, error) {
	panic("implement me")
}
