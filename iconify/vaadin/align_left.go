package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0h11v3H0V0zm0 4h15v3H0V4zm0 4h13v3H0V8zm0 4h16v3H0v-3z"/>`),
		g.Group(children),
	)
}