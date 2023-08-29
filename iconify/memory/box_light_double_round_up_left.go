package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxLightDoubleRoundUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 10h-3v1h-2v1h-2v1h-1v1h-1v1h-1v2h-1v2h-1v3H8v-4h1v-2h1v-2h1v-1h1v-1h1v-1h1v-1h2V9h2V8h4v2M12 22v-2h1v-3h1v-1h1v-1h1v-1h1v-1h3v-1h2v2h-1v1h-3v1h-1v1h-1v1h-1v3h-1v1h-2Z"/>`),
		g.Group(children),
	)
}