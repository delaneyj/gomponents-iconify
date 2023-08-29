package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinusSquareO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 7h8v2H4V7z"/><path fill="currentColor" d="M15 1H1v14h14V1zm-1 13H2V2h12v12z"/>`),
		g.Group(children),
	)
}