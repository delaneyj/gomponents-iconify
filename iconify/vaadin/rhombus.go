package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rhombus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0L0 8l8 8l8-8l-8-8zM2 8l6-6l6 6l-6 6l-6-6z"/>`),
		g.Group(children),
	)
}