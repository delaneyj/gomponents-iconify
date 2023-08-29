package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCatergoriesSofaDecorationFurnitureInteriorDesignCouchSofaDecorate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8" height="7" x="3" y="1.5" rx="1"/><path d="M3 4.5H1.5a1 1 0 0 0-1 1V10a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V5.5a1 1 0 0 0-1-1H11M4 11l-.5 2.5M10 11l.5 2.5"/></g>`),
		g.Group(children),
	)
}