package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.5 19l4.84-7L3.5 5h11c.67 0 1.22.3 1.63.86L20.5 12l-4.37 6.14c-.41.56-.96.86-1.63.86h-11Z"/>`),
		g.Group(children),
	)
}