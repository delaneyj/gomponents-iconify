package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkMultipleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.654 2H10.9c1.16 0 2.1.94 2.1 2.1v8.98a1.8 1.8 0 0 0 1-1.613V4.1A3.1 3.1 0 0 0 10.9 1H6.267a1.8 1.8 0 0 0-1.613 1ZM4.8 3A1.8 1.8 0 0 0 3 4.8v9.7a.5.5 0 0 0 .724.446L7.5 13.06l3.777 1.887A.5.5 0 0 0 12 14.5V4.8A1.8 1.8 0 0 0 10.2 3H4.8Z"/>`),
		g.Group(children),
	)
}