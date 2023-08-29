package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineThicknessTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4.5a.5.5 0 0 1 .5-.5h15a.5.5 0 0 1 0 1h-15a.5.5 0 0 1-.5-.5Zm0 10A1.5 1.5 0 0 1 3.5 13h13a1.5 1.5 0 0 1 0 3h-13A1.5 1.5 0 0 1 2 14.5ZM3 8a1 1 0 0 0 0 2h14a1 1 0 1 0 0-2H3Z"/>`),
		g.Group(children),
	)
}