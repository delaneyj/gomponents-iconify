package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TargetSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM4.5 8a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0ZM8 6a2 2 0 1 0 0 4a2 2 0 0 0 0-4ZM2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8Zm6-4.5a4.5 4.5 0 1 0 0 9a4.5 4.5 0 0 0 0-9Z"/>`),
		g.Group(children),
	)
}