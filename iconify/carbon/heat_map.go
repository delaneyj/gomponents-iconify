package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeatMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 18h2v2h-2zm8-2h2v2h-2zM14 6h2v2h-2zm2 16h-4v-6a2.002 2.002 0 0 0-2-2H4a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h6v4a2.002 2.002 0 0 0 2 2h4a2.002 2.002 0 0 0 2-2v-4a2.002 2.002 0 0 0-2-2zM4 22v-6h6v6zm8 6v-4h4v4zm16 2h-4a2.002 2.002 0 0 1-2-2v-4a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zm-4-6v4h4v-4zm4-22h-6a2.002 2.002 0 0 0-2 2v6h-2a2.002 2.002 0 0 0-2 2v2a2.002 2.002 0 0 0 2 2h2a2.002 2.002 0 0 0 2-2v-2h6a2.002 2.002 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2zM18 14v-2h2v2zm4-4V4h6v6zM8 10H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zM4 4v4h4V4z"/>`),
		g.Group(children),
	)
}