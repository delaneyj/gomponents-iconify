package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailPackage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M4 42h40V18H4v24Z"/><path stroke-linecap="round" d="m4 18l20 15l20-15"/><path stroke-linecap="round" d="M24 18H4v15m40 0V18H24m-12-6h24M16 6h16"/></g>`),
		g.Group(children),
	)
}