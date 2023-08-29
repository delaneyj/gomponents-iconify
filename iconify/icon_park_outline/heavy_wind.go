package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeavyWind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 19.885C4 12.217 10.105 6 17.636 6c6.297 0 11.598 4.346 13.166 10.253a8.921 8.921 0 0 1 4.107-.996c5.02 0 9.091 4.144 9.091 9.257c0 3.642-2.066 6.793-5.07 8.304c-.25.126-.53.182-.81.182H15m0 0h-3a4 4 0 0 0-4 4v0a4 4 0 0 0 4 4h3"/><path d="M22 18h-3a4 4 0 0 0-4 4v0a4 4 0 0 0 4 4h3m-3 0h13"/></g>`),
		g.Group(children),
	)
}