package errs

type Err string

func (e Err) Error() string {
	return string(e)
}

const (
	ServerNA           = Err("server_not_available")
	BadToken           = Err("bad_token")
	NotCorrectPlatform = Err("not_correct_platform")
	BadUserIds         = Err("bad_usr_ids")
	NotAuthorized      = Err("not_authorized")
	JSONParseErr       = Err("json_parse_err")
)
