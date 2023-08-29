package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WingsuitFlying(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-miterlimit="2" stroke-width="4"><path d="M24 14a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 20c-8.71 0-16 7.28-16 16h7v8h18v-8h7c.01-8.71-7.29-16-16-16Zm-9 16V23m18 13V23m-9 21V20"/></g>`),
		g.Group(children),
	)
}