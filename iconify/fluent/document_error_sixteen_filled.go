package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentErrorSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 1h3v3.5A1.5 1.5 0 0 0 10.5 6H14v6.996a2 2 0 0 1-2 2H8.666A5.5 5.5 0 0 0 4 5.205V3a2 2 0 0 1 2-2Zm4.5 4h3.497l-3.989-4H10v3.5a.5.5 0 0 0 .5.5Zm-.5 5.498a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-4.5-2.5a.5.5 0 0 0-.5.5v2a.5.5 0 1 0 1 0v-2a.5.5 0 0 0-.5-.5Zm0 5.125a.625.625 0 1 0 0-1.25a.625.625 0 0 0 0 1.25Z"/>`),
		g.Group(children),
	)
}