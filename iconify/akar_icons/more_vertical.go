package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="4" r="1" transform="rotate(90 12 4)"/><circle cx="12" cy="12" r="1" transform="rotate(90 12 12)"/><circle cx="12" cy="20" r="1" transform="rotate(90 12 20)"/></g>`),
		g.Group(children),
	)
}