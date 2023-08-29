package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandleC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20"/><path d="M32 17.618c-.898-1.83-3.593-5.031-8.983-4.574c-5.39.458-9.433 5.49-8.983 11.893c.45 6.404 5.39 10.063 9.881 10.063C29.305 35 32 30.609 32 30.609"/></g>`),
		g.Group(children),
	)
}