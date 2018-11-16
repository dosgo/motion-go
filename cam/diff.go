package cam

import (
	"flag"
	"fmt"
	"log"
	"os"
	"io"
	"image"
	"strings"
	"net/http"
	"github.com/crhym3/imgdiff"
)
type thresholdVar struct {
	value   float64
	percent bool
}
var threshold = thresholdVar{value: 100}


func diff(up string,next string){


	upimg := readImage(up)
	nextimg := readImage(next)
	res, n, err := newDiffer().Compare(upimg, nextimg)
	if err != nil {
		log.Fatal(err)
	}
	np := float64(n) / float64(res.Bounds().Dx()*res.Bounds().Dy())
	if threshold.percent && !(np > threshold.value) || !(float64(n) > threshold.value) {
		return
	}



	//writeImage(*output, *outputFmt, res)
}


func newDiffer() imgdiff.Differ {
	switch *algorithm {
	case "binary":
		return imgdiff.NewBinary()
	case "perceptual":
		return imgdiff.NewPerceptual(*gamma, *lum, *fov, *cf, *nocolor)
	}
	log.Fatalf("unsupported diff algorithm: %s", *algorithm)
	return nil
}


func readImage(p string) image.Image {
	r := open(p)
	defer r.Close()
	img, _, err := image.Decode(r)
	if err != nil {
		log.Fatalf("%s: %v", p, err)
	}
	return img
}
func open(p string) io.ReadCloser {
	if strings.HasPrefix(p, "http://") || strings.HasPrefix(p, "https://") {
		res, err := http.Get(p)
		if err != nil {
			log.Fatal(err)
		}
		return res.Body
	}
	f, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	return f
}