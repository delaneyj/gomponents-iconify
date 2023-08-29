package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatterySeventyFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 8h2v6H5V8m3 0h2v6H8V8m10-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5h15m-1 2H4v8h13V7m-6 1h2v6h-2V8Z"/>`),
		g.Group(children),
	)
}