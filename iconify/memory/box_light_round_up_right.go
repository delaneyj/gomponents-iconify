package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxLightRoundUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 10h4v1h3v1h1v1h1v1h1v1h1v3h1v4h-2v-3H9v-3H8v-1H7v-1H6v-1H3v-1H0v-2Z"/>`),
		g.Group(children),
	)
}