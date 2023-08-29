package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShiftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 2L1 9h4v5h6V9h4zm2 6v5H6V8H3.5L8 3.42L12.5 8H10z"/>`),
		g.Group(children),
	)
}