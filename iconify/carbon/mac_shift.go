package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MacShift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 28H11a2.002 2.002 0 0 1-2-2V16H4a1 1 0 0 1-.707-1.707l12-12a1 1 0 0 1 1.414 0l12 12A1 1 0 0 1 28 16h-5v10a2.003 2.003 0 0 1-2 2ZM6.414 14H11v12h10V14h4.586L16 4.414Z"/>`),
		g.Group(children),
	)
}