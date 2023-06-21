package bankimpl

// YBCCB 宜宾市商业银行
type YBCCB struct {
}

func init() {
	RegisterBank("YBCCB", new(YBCCB))
}

// OpenAccount open
func (C *YBCCB) OpenAccount(args *OpenAccountArgs) (*OpenAccountRet, error) {
	panic("implement me")
}
