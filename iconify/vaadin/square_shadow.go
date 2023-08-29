package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareShadow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 2V0H0v14h2v2h14V2h-2zm-1 11H1V1h12v12z"/>`),
		g.Group(children),
	)
}