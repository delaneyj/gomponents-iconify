package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpCheckBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3H3v18h18V3zM10 17l-5-5l1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>`),
		g.Group(children),
	)
}