package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M5.25 20.5h13.498a.75.75 0 0 1 .101 1.493l-.101.007H5.25a.75.75 0 0 1-.102-1.493l.102-.007h13.498H5.25zM6.293 7.29l4.998-4.997a1 1 0 0 1 1.32-.084l.093.083l5.004 4.997a1 1 0 0 1-1.32 1.498l-.094-.083L13 5.413V18a1 1 0 0 1-.884.993L12 19a1 1 0 0 1-.993-.883L10.999 18L11 5.41L7.707 8.704a1 1 0 0 1-1.32.083l-.094-.083a1 1 0 0 1-.083-1.32l.083-.095l4.998-4.996l-4.998 4.996z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}