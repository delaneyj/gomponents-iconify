package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOuterLightDashedRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M20 21v-4h2v4h-2m0-6v-3h2v3h-2m0-5V6h2v4h-2m0-6V1h2v3h-2Z"/>`),
		g.Group(children),
	)
}