package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Indent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0h16v3H0V0zm6 4h10v3H6V4zm0 4h10v3H6V8zm-6 4h16v3H0v-3zm0-7.5v6l4-3z"/>`),
		g.Group(children),
	)
}