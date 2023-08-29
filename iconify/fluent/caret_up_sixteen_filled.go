package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretUpSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.957 10.998a1 1 0 0 1-.821-1.571l2.633-3.784a1.5 1.5 0 0 1 2.462 0l2.633 3.784a1 1 0 0 1-.821 1.571H4.957Z"/>`),
		g.Group(children),
	)
}