package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretSquareDownO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 1H1v14h14V1zm-1 13H2V2h12v12z"/><path fill="currentColor" d="M4 6h8l-4 5z"/>`),
		g.Group(children),
	)
}