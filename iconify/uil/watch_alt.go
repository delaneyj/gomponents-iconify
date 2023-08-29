package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 8.61l-1-5.77A1 1 0 0 0 15 2H9a1 1 0 0 0-1 .84L7 8.61a6 6 0 0 0 0 6.78l1 5.77A1 1 0 0 0 9 22h6a1 1 0 0 0 1-.84l1-5.77a6 6 0 0 0 0-6.78ZM9.85 4h4.3l.44 2.59a6 6 0 0 0-5.18 0Zm4.3 16h-4.3l-.44-2.59a6 6 0 0 0 5.18 0ZM12 16a4 4 0 1 1 4-4a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}