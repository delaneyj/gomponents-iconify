package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoordinateSystem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1l11.258 19.5H.74L12 1ZM4.204 18.5h15.589L11.999 5L4.205 18.5ZM13 9v3.532l3.409 2.84l-1.28 1.536L12 14.302l-3.128 2.606l-1.28-1.536L11 12.532V9h2Z"/>`),
		g.Group(children),
	)
}