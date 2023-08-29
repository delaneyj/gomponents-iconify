package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="20" cy="8" r="1" transform="rotate(-180 20 8)"/><circle cx="20" cy="16" r="1" transform="rotate(-180 20 16)"/><circle cx="12" cy="8" r="1" transform="rotate(-180 12 8)"/><circle cx="12" cy="16" r="1" transform="rotate(-180 12 16)"/><circle cx="4" cy="8" r="1" transform="rotate(-180 4 8)"/><circle cx="4" cy="16" r="1" transform="rotate(-180 4 16)"/></g>`),
		g.Group(children),
	)
}