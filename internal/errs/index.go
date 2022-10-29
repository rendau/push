package errs

import (
	"github.com/rendau/dop/dopErrs"
)

const (
	PlatformRequired   = dopErrs.Err("platform_required")
	BadPlatform        = dopErrs.Err("bad_platform")
	TokenValueRequired = dopErrs.Err("token_value_required")
	BadTokenValue      = dopErrs.Err("bad_token_value")
	UsrIdRequired      = dopErrs.Err("usr_id_required")
)
