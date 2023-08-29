package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.7 8.46a.75.75 0 1 0-1.5 0v6.68L7.58 6.52a.75.75 0 0 0-1.06 1.06l8.62 8.62H8.46a.75.75 0 0 0 0 1.5H17a.75.75 0 0 0 .29-.06a.76.76 0 0 0 .41-.64Z"/>`),
		g.Group(children),
	)
}