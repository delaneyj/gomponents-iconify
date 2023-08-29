package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankCardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M44 18V8H4v10"/><path d="M44 18H4v22h40V18Z"/><path stroke-linecap="round" d="M12 29h2m6 0h2m6 0h2"/></g>`),
		g.Group(children),
	)
}