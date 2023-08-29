package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M29 16a1 1 0 0 1-1 1H6.414l8.293 8.293a1 1 0 0 1-1.414 1.414l-10-10a1 1 0 0 1 0-1.414l10-10a1 1 0 1 1 1.414 1.414L6.414 15H28a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}