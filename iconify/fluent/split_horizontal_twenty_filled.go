package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitHorizontalTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.5 10a.5.5 0 0 0 0-1h-15a.5.5 0 0 0 0 1h15ZM14 2a2 2 0 0 1 2 2v4H4V4a2 2 0 0 1 2-2h8ZM4 11v4a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-4H4Z"/>`),
		g.Group(children),
	)
}