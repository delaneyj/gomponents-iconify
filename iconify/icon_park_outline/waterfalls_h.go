package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterfallsH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M20 6H6v11h14V6Zm22 25H28v11h14V31Zm0-25H28v17h14V6ZM20 25H6v17h14V25Z"/>`),
		g.Group(children),
	)
}