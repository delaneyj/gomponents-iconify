package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15.5 8.516a7.5 7.5 0 1 1-9.462-7.24A1 1 0 0 1 7 0h2a1 1 0 0 1 .962 1.276a7.503 7.503 0 0 1 5.538 7.24zm-3.61-3.905L6.94 7.439L4.11 12.39l4.95-2.828l2.828-4.95z"/>`),
		g.Group(children),
	)
}