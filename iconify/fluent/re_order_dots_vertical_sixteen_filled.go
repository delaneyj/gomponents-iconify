package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReOrderDotsVerticalSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 4.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm1.5 3a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM10.5 5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM12 8a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm-1.5 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}