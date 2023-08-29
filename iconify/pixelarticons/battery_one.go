package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5H2v14h18v-4h2V9h-2V5H4zm14 2v10H4V7h14zM8 9H6v6h2V9z"/>`),
		g.Group(children),
	)
}