package firebase

import (
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

var opt option.ClientOption

func init() {
	if gin.Mode() == gin.ReleaseMode {
		opt = option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_KEYFILE_JSON")))
	} else {
		opt = option.WithCredentialsFile("serviceAccountKey.json")
	}
}
