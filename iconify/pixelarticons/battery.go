package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Battery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5H2v14h18v-4h2V9h-2V5H4zm14 2v10H4V7h14z"/>`),
		g.Group(children),
	)
}