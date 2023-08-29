package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoDimensionalCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M20 6H6v14h14V6Zm0 22H6v14h14V28ZM42 6H28v14h14V6Z"/><path stroke-linecap="round" d="M29 28v14m12-14v14"/></g>`),
		g.Group(children),
	)
}