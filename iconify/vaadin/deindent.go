package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deindent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 10.5v-6l-4 3zM0 0h16v3H0V0zm6 4h10v3H6V4zm0 4h10v3H6V8zm-6 4h16v3H0v-3z"/>`),
		g.Group(children),
	)
}