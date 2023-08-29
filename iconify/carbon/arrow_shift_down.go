package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowShiftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 6v10h5a1 1 0 0 1 .707 1.707l-12 12a1 1 0 0 1-1.414 0l-12-12A1 1 0 0 1 4 16h5V6a2.002 2.002 0 0 1 2-2h10a2.003 2.003 0 0 1 2 2Zm-7 21.586L25.586 18H21V6H11v12H6.414Z"/>`),
		g.Group(children),
	)
}