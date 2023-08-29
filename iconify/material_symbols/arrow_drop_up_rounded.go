package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDropUpRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.425 14q-.675 0-.938-.613T8.7 12.3l2.6-2.6q.15-.15.325-.225T12 9.4q.2 0 .375.075t.325.225l2.6 2.6q.475.475.212 1.088t-.937.612h-5.15Z"/>`),
		g.Group(children),
	)
}