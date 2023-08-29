package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1453 499l-90 90l-338-337l-1 1796H896l1-1799l-340 340l-90-90L960 6l493 493z"/>`),
		g.Group(children),
	)
}