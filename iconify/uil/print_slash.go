package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 10a1 1 0 1 0 1 1a1 1 0 0 0-1-1ZM3.71 2.29a1 1 0 0 0-1.42 1.42L4.62 6A3 3 0 0 0 2 9v6a3 3 0 0 0 3 3h1v3a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-1.59l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM6 15v1H5a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h1.59l6 6H7a1 1 0 0 0-1 1Zm10 5H8v-4h6.59L16 17.41Zm3-14h-1V3a1 1 0 0 0-1-1H8.66a1 1 0 0 0 0 2H16v2h-3.34a1 1 0 0 0 0 2H19a1 1 0 0 1 1 1v6a.37.37 0 0 1 0 .11a1 1 0 0 0 .78 1.18h.2a1 1 0 0 0 1-.8A2.84 2.84 0 0 0 22 15V9a3 3 0 0 0-3-3Z"/>`),
		g.Group(children),
	)
}