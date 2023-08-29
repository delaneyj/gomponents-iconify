package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M4 8h40v8s-10 2-10 8s10 8 10 8v8H4V8Z"/><path stroke-linecap="round" d="M44 24h-2"/></g>`),
		g.Group(children),
	)
}