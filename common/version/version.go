package version

import (
	"fmt"
	"net/http"
)

var (
	BuildDate    string
	BuildVersion string
	RevisionDate string
)

func init() {
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		info := fmt.Sprintf(`{"build_date":"%s", "build_version":"%s", "revision_date":"%s"}`,
			BuildDate,
			BuildVersion,
			RevisionDate,
		)
		w.WriteHeader(200)
		w.Write([]byte(info))
	})
}
