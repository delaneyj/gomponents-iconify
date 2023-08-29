package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7 1v2l1 1V2h7v12H8v-2l-1 1v2h9V1z"/><path fill="currentColor" d="M10 8L5 4v2H0v4h5v2z"/>`),
		g.Group(children),
	)
}