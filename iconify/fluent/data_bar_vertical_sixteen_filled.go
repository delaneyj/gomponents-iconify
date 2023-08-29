package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataBarVerticalSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 2A1.5 1.5 0 0 0 2 3.5v9a1.5 1.5 0 0 0 3 0v-9A1.5 1.5 0 0 0 3.5 2Zm4 3A1.5 1.5 0 0 0 6 6.5v6a1.5 1.5 0 0 0 3 0v-6A1.5 1.5 0 0 0 7.5 5Zm4 3A1.5 1.5 0 0 0 10 9.5v3a1.5 1.5 0 0 0 3 0v-3A1.5 1.5 0 0 0 11.5 8Z"/>`),
		g.Group(children),
	)
}