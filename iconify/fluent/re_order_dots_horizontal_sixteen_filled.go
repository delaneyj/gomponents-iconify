package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReOrderDotsHorizontalSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 5.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0Zm-4.5 0a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0ZM3.5 7a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm7.5 3.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0ZM8 12a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm-6-1.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0Z"/>`),
		g.Group(children),
	)
}