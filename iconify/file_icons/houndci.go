package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Houndci(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 0h512v224.92C490 114 348 36 358.361 75.615c9.933 23.588 9.933 85.033 9.933 85.033l-100.01 38.146l-15 55.482l-122.398 32.27s8.722 62.43 43.903 73.341c22.566 6.999 170.482 23.316 170.482 23.316l-28.08 123.778H0V0z"/>`),
		g.Group(children),
	)
}