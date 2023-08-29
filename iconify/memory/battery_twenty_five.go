package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryTwentyFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M7 8v6H5V8h2m11-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5h15m-1 2H4v8h13V7Z"/>`),
		g.Group(children),
	)
}