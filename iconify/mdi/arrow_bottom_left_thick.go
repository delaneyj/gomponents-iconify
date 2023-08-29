package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomLeftThick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.5 5.69l2.81 2.81l-6.37 6.39h4.95v3.42H5.69V7.11h3.43v4.95l6.38-6.37Z"/>`),
		g.Group(children),
	)
}