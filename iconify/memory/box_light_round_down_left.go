package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxLightRoundDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 12h-4v-1h-3v-1h-1V9h-1V8h-1V7h-1V4h-1V0h2v3h1v3h1v1h1v1h1v1h3v1h3v2Z"/>`),
		g.Group(children),
	)
}