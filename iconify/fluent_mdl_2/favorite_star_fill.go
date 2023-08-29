package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FavoriteStarFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1416 1254l248 794l-640-492l-640 492l248-794L0 768h784L1024 0l240 768h784l-632 486z"/>`),
		g.Group(children),
	)
}