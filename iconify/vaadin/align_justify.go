package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignJustify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0h16v3H0V0zm0 4h16v3H0V4zm0 8h16v3H0v-3zm0-4h16v3H0V8z"/>`),
		g.Group(children),
	)
}