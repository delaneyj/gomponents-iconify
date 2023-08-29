package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleMultipleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6.5A4.5 4.5 0 0 1 6.5 2h11a4.5 4.5 0 1 1 0 9h-11A4.5 4.5 0 0 1 2 6.5Zm4.5 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 4.5a4.5 4.5 0 1 0 0 9h11a4.5 4.5 0 1 0 0-9h-11Zm13 4.5a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/>`),
		g.Group(children),
	)
}