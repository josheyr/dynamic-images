package generate

import (
	"encoding/base64"
	"net/http"
)

const transPngPixelB64 = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+P+/HgAFhAJ/wlseKgAAAABJRU5ErkJggg=="

func HandlePixel(w http.ResponseWriter, r *http.Request) {
	// render 1 pixel png
	// return png

	// decode base64
	decoded, err := base64.StdEncoding.DecodeString(transPngPixelB64)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(decoded)
}
