package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Delete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M2 1L0 4l2 3h6V1H2zm1.5.78L5 3.28l1.5-1.5l.72.72L5.72 4l1.5 1.5l-.72.72L5 4.72l-1.5 1.5l-.72-.72L4.28 4l-1.5-1.5l.72-.72z"/>`),
		g.Group(children),
	)
}