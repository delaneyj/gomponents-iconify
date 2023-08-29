package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackSquareButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#333" d="M2 2h60v60H2z"/><path fill="#d0d0d0" d="M10 10h44v44H10z"/>`),
		g.Group(children),
	)
}