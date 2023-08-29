package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOuterLightRoundDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 22v-2h5v-1h4v-1h2v-1h2v-1h1v-1h1v-1h1v-1h1v-2h1V9h1V5h1V0h2v6h-1v4h-1v2h-1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-2v1h-2v1H6v1H0Z"/>`),
		g.Group(children),
	)
}