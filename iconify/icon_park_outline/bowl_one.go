package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowlOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M24 39c9.389 0 17-7.059 17-17H7c0 9.941 7.611 17 17 17Z"/><path stroke-linejoin="round" d="m18 38l-2 6h16l-2-6"/><path d="M12 10v4m24-4v4M24 5v9"/></g>`),
		g.Group(children),
	)
}