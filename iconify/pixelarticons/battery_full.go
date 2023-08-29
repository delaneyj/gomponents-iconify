package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 5H2v14h18v-4h2V9h-2V5h-2zm0 2v10H4V7h14zM8 9H6v6h2V9zm2 0h2v6h-2V9zm6 0h-2v6h2V9z"/>`),
		g.Group(children),
	)
}