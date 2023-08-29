package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sweater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 37H6V24l5-15l8-5h10l9 5l4 15v13h-8v7H14v-7Zm20-9v9m-20-9v9"/>`),
		g.Group(children),
	)
}