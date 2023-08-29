package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretSquareRightO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 1H1v14h14V1zm-1 13H2V2h12v12z"/><path fill="currentColor" d="M5.9 12V4l5 4z"/>`),
		g.Group(children),
	)
}