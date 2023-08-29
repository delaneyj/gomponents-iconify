package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPinTimeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.95 15.95a7 7 0 1 0-9.9 0L12 20.9l4.95-4.95ZM12 23.728l-6.364-6.364a9 9 0 1 1 12.728 0L12 23.728ZM13 11h4v2h-6V6h2v5Z"/>`),
		g.Group(children),
	)
}