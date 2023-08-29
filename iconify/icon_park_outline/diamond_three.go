package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="4" d="M5.414 22.586L22.586 5.414a2 2 0 0 1 2.828 0l17.172 17.172a2 2 0 0 1 0 2.828L25.414 42.586a2 2 0 0 1-2.828 0L5.414 25.414a2 2 0 0 1 0-2.828Z"/>`),
		g.Group(children),
	)
}