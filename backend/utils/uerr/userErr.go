package uerr

type UserErrorData struct {
	endUserSafeMsg string
	returnCode     int
	innerErr       error
	shouldLog      bool
}

func (s *UserErrorData) Error() string {
	return s.endUserSafeMsg
}

func (s *UserErrorData) Unwrap() error {
	return s.innerErr
}

func (s *UserErrorData) UserMsgAndCode() (string, int) {
	return s.endUserSafeMsg, s.returnCode
}

func (s *UserErrorData) UserCode() int {
	return s.returnCode
}

func (s *UserErrorData) ShouldLog() bool {
	return s.shouldLog
}

func UErr(endUserSafeMsg string, returnCode int) error {
	ret := &UserErrorData{
		endUserSafeMsg: endUserSafeMsg,
		returnCode:     returnCode,
	}
	return ret
}

func UErrLog(endUserSafeMsg string, returnCode int, innerErr error) error {
	ret := &UserErrorData{
		endUserSafeMsg: endUserSafeMsg,
		returnCode:     returnCode,
		innerErr:       innerErr,
		shouldLog:      true,
	}
	return ret
}

func UErrLogHash(endUserSafeMsg string, returnCode int, innerErr error) error {
	endUserSafeMsg = endUserSafeMsg + " " + HashError(innerErr)
	return UErrLog(endUserSafeMsg, returnCode, innerErr)
}
