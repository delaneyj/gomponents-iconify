package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 5h15v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5m1 2v8h13V7H4Z"/>`),
		g.Group(children),
	)
}