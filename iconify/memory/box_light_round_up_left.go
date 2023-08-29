package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxLightRoundUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 22v-4h1v-3h1v-1h1v-1h1v-1h1v-1h3v-1h4v2h-3v1h-3v1h-1v1h-1v1h-1v3h-1v3h-2Z"/>`),
		g.Group(children),
	)
}