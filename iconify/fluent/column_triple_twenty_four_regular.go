package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColumnTripleTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 21A2.5 2.5 0 0 1 2 18.5v-13A2.5 2.5 0 0 1 4.5 3h1A2.5 2.5 0 0 1 8 5.5v13A2.5 2.5 0 0 1 5.5 21h-1Zm-1-2.5a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-13a1 1 0 0 0-1-1h-1a1 1 0 0 0-1 1v13Zm8 2.5A2.5 2.5 0 0 1 9 18.5v-13A2.5 2.5 0 0 1 11.5 3h1A2.5 2.5 0 0 1 15 5.5v13a2.5 2.5 0 0 1-2.5 2.5h-1Zm-1-2.5a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-13a1 1 0 0 0-1-1h-1a1 1 0 0 0-1 1v13Zm5.5 0a2.5 2.5 0 0 0 2.5 2.5h1a2.5 2.5 0 0 0 2.5-2.5v-13A2.5 2.5 0 0 0 19.5 3h-1A2.5 2.5 0 0 0 16 5.5v13Zm2.5 1a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v13a1 1 0 0 1-1 1h-1Z"/>`),
		g.Group(children),
	)
}