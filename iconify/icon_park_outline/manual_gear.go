package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManualGear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 12v12H8m16-12v24M8 12v24"/><path d="M44 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0ZM28 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0ZM12 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm16 32a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm-16 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm28 4a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g>`),
		g.Group(children),
	)
}