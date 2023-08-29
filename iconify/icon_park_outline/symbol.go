package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Symbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 40.003h14.004C11.334 35.636 8 30.301 8 24c0-9.453 7.017-16 16.008-16C33 8 40 15 40 24c0 6-3.331 11.334-9.993 16.003H44"/>`),
		g.Group(children),
	)
}