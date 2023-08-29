package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumericPositiveOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 7v2h2v8h2V7h-4m-2 6H9v2H7v-2H5v-2h2V9h2v2h2v2Z"/>`),
		g.Group(children),
	)
}