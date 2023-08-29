package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxLightDashedUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 10h2v2H0v-2m4 0h3v2H4v-2m5 0h3v3h-2v-1H9v-2m1 12v-2h2v2h-2m0-4v-3h2v3h-2Z"/>`),
		g.Group(children),
	)
}