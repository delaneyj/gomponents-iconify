package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorPointSelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 20l-5 2l5-11l5 11l-5-2M8 2h8v3h6v2h-6v3H8V7H2V5h6V2m2 2v4h4V4h-4Z"/>`),
		g.Group(children),
	)
}