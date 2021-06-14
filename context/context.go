package context

import "os"

var (
	hueClientId     = ""
	hueClientSecret = ""
	hueAppId        = ""
	Version         = ""
)

func init() {
	os.Setenv("HUE_CLIENT_ID", hueClientId)
	os.Setenv("HUE_CLIENT_SECRET", hueClientSecret)
	os.Setenv("HUE_APP_ID", hueAppId)
}
