package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOuterLightDashedDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M18 22v-2h2v-2h2v4h-4m2-6v-4h2v4h-2m0-6V6h2v4h-2m0-6V1h2v3h-2m-4 18h-4v-2h4v2m-6 0H7v-2h3v2m-5 0H1v-2h4v2Z"/>`),
		g.Group(children),
	)
}