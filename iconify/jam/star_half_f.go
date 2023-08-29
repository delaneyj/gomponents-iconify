package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalfF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.989 0L10 .024V15.76l-.011-.006L3.815 19l1.179-6.874L0 7.257l6.902-1.003z"/>`),
		g.Group(children),
	)
}