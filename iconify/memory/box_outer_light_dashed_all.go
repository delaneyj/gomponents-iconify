package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOuterLightDashedAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 0v2H2v2H0V0h4M2 6v4H0V6h2m0 6v4H0v-4h2m0 6v2h2v2H0v-4h2M6 0h4v2H6V0m6 0h4v2h-4V0m6 0h4v4h-2V2h-2V0m0 22v-2h2v-2h2v4h-4m-2 0h-4v-2h4v2m-6 0H6v-2h4v2M20 6h2v4h-2V6m0 6h2v4h-2v-4Z"/>`),
		g.Group(children),
	)
}