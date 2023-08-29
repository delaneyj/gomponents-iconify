package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CafeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 64H80v254.34a19.83 19.83 0 0 0 5.86 14.14l29.65 29.66a19.87 19.87 0 0 0 14.15 5.86h204.68a19.87 19.87 0 0 0 14.15-5.86l29.65-29.66a19.83 19.83 0 0 0 5.86-14.14V192h32a16 16 0 0 0 11.31-4.69l32-32A16 16 0 0 0 464 144V80a16 16 0 0 0-16-16Zm-16 73.37L409.37 160H384V96h48ZM48 400h368v32H48z"/>`),
		g.Group(children),
	)
}