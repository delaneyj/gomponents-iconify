package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 0v16h12V0H2zm11 15H9v-3H7v3H3V1h10v14z"/><path fill="currentColor" d="M4 9h2v2H4V9zm3 0h2v2H7V9zm3 0h2v2h-2V9zM4 6h2v2H4V6zm3 0h2v2H7V6zm3 0h2v2h-2V6zM4 3h2v2H4V3zm3 0h2v2H7V3zm3 0h2v2h-2V3z"/>`),
		g.Group(children),
	)
}