package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOuterLightDashedDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 20h4v2H1v-2m6 0h3v2H7v-2m5 0h4v2h-4v-2m6 0h3v2h-3v-2Z"/>`),
		g.Group(children),
	)
}