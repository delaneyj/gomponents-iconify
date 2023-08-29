package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16zm1.586-7H6a1 1 0 0 1 0-2h5.586L9.05 6.464a1 1 0 1 1 1.414-1.414l4.243 4.243a.997.997 0 0 1 0 1.414l-4.243 4.243a1 1 0 1 1-1.414-1.414L11.586 11z"/>`),
		g.Group(children),
	)
}