package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sliders(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7 0h2v3H7V0zM6 4v3h1v9h2V7h1V4zM2 0h2v8H2V0zM1 9v3h1v4h2v-4h1V9zm11-9h2v10h-2V0zm-1 11v3h1v2h2v-2h1v-3z"/>`),
		g.Group(children),
	)
}