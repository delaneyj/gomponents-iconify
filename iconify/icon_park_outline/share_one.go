package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M35 16a5 5 0 1 0 0-10a5 5 0 0 0 0 10ZM13 29a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" d="m30 13.575l-12.66 7.67m-.002 5.319l13.34 7.883"/><path d="M35 32a5 5 0 1 1 0 10a5 5 0 0 1 0-10Z"/></g>`),
		g.Group(children),
	)
}