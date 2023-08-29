package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourK(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7h2v4h4V7h2v10H9v-4H3V7zm10 0h2v4h2v2h-2v4h-2V7zm6 8h-2v-2h2v2zm0 0h2v2h-2v-2zm0-6h-2v2h2V9zm0 0V7h2v2h-2z"/>`),
		g.Group(children),
	)
}