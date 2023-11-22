package utils

type UserErr struct {
	endUserSafeMsg string
	returnCode     int
	innerErr       error
	shouldLog      bool
}

func (s *UserErr) Error() string {
	return s.endUserSafeMsg
}

func (s *UserErr) Unwrap() error {
	return s.innerErr
}

func (s *UserErr) UserMsgAndCode() (string, int) {
	return s.endUserSafeMsg, s.returnCode
}

func (s *UserErr) UserCode() int {
	return s.returnCode
}

func (s *UserErr) ShouldLog() bool {
	return s.shouldLog
}

func NewUserErr(endUserSafeMsg string, returnCode int) error {
	ret := &UserErr{
		endUserSafeMsg: endUserSafeMsg,
		returnCode:     returnCode,
	}
	return ret
}

func NewUserErrLog(endUserSafeMsg string, returnCode int, innerErr error) error {
	ret := &UserErr{
		endUserSafeMsg: endUserSafeMsg,
		returnCode:     returnCode,
		innerErr:       innerErr,
		shouldLog:      true,
	}
	return ret
}

func NewUserErrLogHash(endUserSafeMsg string, returnCode int, innerErr error) error {
	endUserSafeMsg = endUserSafeMsg + " " + HashError(innerErr)
	return NewUserErrLog(endUserSafeMsg, returnCode, innerErr)
}
