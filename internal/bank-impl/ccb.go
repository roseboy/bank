package bankimpl

// CCB 中国建设银行
type CCB struct {
}

func init() {
	RegisterBank("CCB", new(CCB))
}

// OpenAccount open
func (C *CCB) OpenAccount(args *OpenAccountArgs) (*OpenAccountRet, error) {
	return &OpenAccountRet{AccountNum: "unimplemented"}, nil
}
