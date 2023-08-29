package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelLeftHeaderAddTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v4.257a5.503 5.503 0 0 0-1-.657V8H8.5v7h.522a5.5 5.5 0 0 0 .185 1H5a3 3 0 0 1-3-3V6Zm15 0a2 2 0 0 0-2-2H8.5v3H17V6Zm2 8.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-4-2a.5.5 0 0 0-1 0V14h-1.5a.5.5 0 0 0 0 1H14v1.5a.5.5 0 0 0 1 0V15h1.5a.5.5 0 0 0 0-1H15v-1.5Z"/>`),
		g.Group(children),
	)
}