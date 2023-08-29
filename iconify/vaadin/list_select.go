package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListSelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1 0h12v2H1V0zm0 8h13v2H1V8zm0 3h11v2H1v-2zm0 3h14v2H1v-2zM0 3v4h16V3H0zm11 3H1V4h10v2z"/>`),
		g.Group(children),
	)
}