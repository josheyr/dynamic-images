package generate

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"github.com/josheyr/svg2png"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GenerateDynamicImages(w http.ResponseWriter, r *http.Request) {
	// get svg from uploaded content
	// convert to png
	// return png

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return
	}

	val := r.FormValue("svg")

	csvVal, _, err := r.FormFile("csv")
	// loop csv with reader
	csvReader := csv.NewReader(csvVal)

	// make zip
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	headers, err := csvReader.Read()
	if err != nil {
		return
	}

	// loop records
	for i := 0; ; i++ {
		record, err := csvReader.Read()
		if err != nil {
			break
		}

		// replace everything within {{}} with the value from the csv
		// eg. {{hello}}
		for j, header := range headers {
			val = strings.ReplaceAll(val, "{{"+header+"}}", record[j])

			expected := "http://localhost:8080/emptypixel/" + url.PathEscape("{{"+header+"}}") + ".png"
			val = strings.ReplaceAll(val, expected, record[j])

			val = strings.ReplaceAll(val, "width=\"1\" height=\"1\"", "")
		}

		// replace values in svg
		// convert to png
		png, err := svg2png.SvgToPng(val, 600, 600)
		if err != nil {
			return
		}

		// write png to zip
		create, err := zipWriter.Create(strconv.Itoa(i) + ".png")
		if err != nil {
			return
		}

		_, err = create.Write(png)
		if err != nil {
			return
		}

		val = r.FormValue("svg")
	}

	err = zipWriter.Close()
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "image/zip")
	// return png
	_, err = w.Write(buf.Bytes())
	if err != nil {
		return
	}
}
