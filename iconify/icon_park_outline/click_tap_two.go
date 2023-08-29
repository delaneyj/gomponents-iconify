package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClickTapTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M11 44V15a4 4 0 0 1 8 0v20l23 5v4"/><path d="M11 25.25v0C6.903 23.65 4 19.664 4 15C4 8.925 8.925 4 15 4s11 4.925 11 11c0 4.664-2.903 8.65-7 10.25"/></g>`),
		g.Group(children),
	)
}