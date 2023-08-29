package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChildWithPacifier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M24 42c9.389 0 17-7.611 17-17S33.389 8 24 8S7 15.611 7 25s7.611 17 17 17Z"/><path stroke-linecap="round" d="m33 21l-2 1l-2-1m-10 0l-2 1l-2-1m9 17v6m0-36c-.25-1-2-4-6-4m6 4c.083-1 .6-3.2 2-4m5.975 36c0-6 0-9-7.975-9c-8 0-7.976 3-7.976 9M5 23v4m38-4v4"/></g>`),
		g.Group(children),
	)
}