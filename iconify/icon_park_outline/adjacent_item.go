package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjacentItem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="M14 29h28v12H14V29Zm0-22h28v12H14V7Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 13v6h28V7H14v6Zm0 0H6v22h8m0 0v6h28V29H14v6Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 13H6v22h8m0-6h28v12H14V29Zm0-22h28v12H14V7Z"/></g>`),
		g.Group(children),
	)
}