package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InHomeMode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.2 22L10 17.8l1.4-1.4l2.8 2.8l5.9-5.9l1.4 1.4l-7.3 7.3ZM4 20V10H1l11-9l11 9h-3v.575l-5.8 5.8l-2.8-2.8L7.175 17.8l2.2 2.2H4Z"/>`),
		g.Group(children),
	)
}