package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TShirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m9 9l9-5h12l9 5l4 15l-8 6v14H13V30l-8-6L9 9Zm4 22v-7m22 7v-7"/>`),
		g.Group(children),
	)
}