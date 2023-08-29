package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7 0h2v16H7V0zM3 5L0 8l3 3V9h3V7H3zm13 3l-3-3v2h-3v2h3v2z"/>`),
		g.Group(children),
	)
}