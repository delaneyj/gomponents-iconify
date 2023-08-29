package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowStepInDiagonalDownLeftSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.854 2.146a.5.5 0 0 1 0 .708L8.707 8H12.5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 1 0v3.793l5.146-5.147a.5.5 0 0 1 .708 0ZM6 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm-1 0a1 1 0 1 0-2 0a1 1 0 0 0 2 0Z"/>`),
		g.Group(children),
	)
}