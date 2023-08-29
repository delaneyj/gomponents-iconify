package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 4V1H0v14h9v-3H8v2H1V2h7v2z"/><path fill="currentColor" d="m16 8l-5-4v2H6v4h5v2z"/>`),
		g.Group(children),
	)
}