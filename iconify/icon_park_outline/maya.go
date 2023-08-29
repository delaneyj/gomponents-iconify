package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maya(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M4 34h40v8H4zm3-7h34v7H7zm3-6h28v6H10zm3-6h22v6H13zm7-8h8v8h-8z"/><path d="m20 15l-6 27m14-27l6 27"/></g>`),
		g.Group(children),
	)
}