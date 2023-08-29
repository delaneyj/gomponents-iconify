package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassDiscoverFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 22C7.477 22 3 17.523 3 12S7.477 2 13 2s10 4.477 10 10s-4.477 10-10 10ZM8 11.5l4 1.5l1.5 4.002L17 8l-9 3.5Z"/>`),
		g.Group(children),
	)
}