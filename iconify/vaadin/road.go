package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Road(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 11v4h7L12 1H9v3H7V1H4L0 15h7v-4h2zM7 6h2v3H7V6z"/>`),
		g.Group(children),
	)
}