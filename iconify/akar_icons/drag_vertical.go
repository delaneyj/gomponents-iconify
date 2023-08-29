package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="8" cy="4" r="1" transform="rotate(90 8 4)"/><circle cx="16" cy="4" r="1" transform="rotate(90 16 4)"/><circle cx="8" cy="12" r="1" transform="rotate(90 8 12)"/><circle cx="16" cy="12" r="1" transform="rotate(90 16 12)"/><circle cx="8" cy="20" r="1" transform="rotate(90 8 20)"/><circle cx="16" cy="20" r="1" transform="rotate(90 16 20)"/></g>`),
		g.Group(children),
	)
}