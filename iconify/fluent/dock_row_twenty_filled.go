package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DockRowTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 4a1.5 1.5 0 0 0-1.5 1.5v2A1.5 1.5 0 0 0 3 9h2a1.5 1.5 0 0 0 1.5-1.5v-2A1.5 1.5 0 0 0 5 4H3Zm0 6a1.5 1.5 0 0 0-1.5 1.5v2A1.5 1.5 0 0 0 3 15h2a1.5 1.5 0 0 0 1.5-1.5v-2A1.5 1.5 0 0 0 5 10H3Zm4.5-4.5A1.5 1.5 0 0 1 9 4h2a1.5 1.5 0 0 1 1.5 1.5v2A1.5 1.5 0 0 1 11 9H9a1.5 1.5 0 0 1-1.5-1.5v-2ZM9 10a1.5 1.5 0 0 0-1.5 1.5v2A1.5 1.5 0 0 0 9 15h2a1.5 1.5 0 0 0 1.5-1.5v-2A1.5 1.5 0 0 0 11 10H9Zm4.5-4.5A1.5 1.5 0 0 1 15 4h2a1.5 1.5 0 0 1 1.5 1.5v2A1.5 1.5 0 0 1 17 9h-2a1.5 1.5 0 0 1-1.5-1.5v-2ZM15 10a1.5 1.5 0 0 0-1.5 1.5v2A1.5 1.5 0 0 0 15 15h2a1.5 1.5 0 0 0 1.5-1.5v-2A1.5 1.5 0 0 0 17 10h-2Z"/>`),
		g.Group(children),
	)
}