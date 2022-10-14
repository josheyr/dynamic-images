package generate

import (
	"github.com/josheyr/svg2png"
	"net/http"
)

func GenerateSingleImage(w http.ResponseWriter, r *http.Request) {
	// get svg from uploaded content
	// convert to png
	// return png

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return
	}

	val := r.FormValue("svg")

	png, err := svg2png.SvgToPng(val, 600, 600)

	w.Header().Set("Content-Type", "image/png")
	// return png
	_, err = w.Write(png)
	if err != nil {
		return
	}
}
