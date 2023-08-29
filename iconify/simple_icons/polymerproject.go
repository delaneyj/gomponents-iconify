package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Polymerproject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.4 3.686L7.2 16.16L4.8 12l4.8-8.314H4.8L0 12l2.4 4.159l2.4 4.155h4.8l7.2-12.469L19.2 12l-4.8 8.314h4.8l2.4-4.155L24 12l-2.4-4.155l-2.4-4.159Z"/>`),
		g.Group(children),
	)
}