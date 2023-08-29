package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReOrderDotsVerticalTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM8 14.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM13.5 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM15 9.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM13.5 16a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}