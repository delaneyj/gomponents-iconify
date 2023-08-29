package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraserFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 18.997h7v2h-9l-3.998.002l-6.487-6.487a1 1 0 0 1 0-1.415L12.12 2.491a1 1 0 0 1 1.414 0l7.779 7.778a1 1 0 0 1 0 1.414L14 18.997Zm1.657-4.486l3.535-3.535l-6.364-6.364l-3.535 3.536l6.364 6.364Z"/>`),
		g.Group(children),
	)
}