package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeMedium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M16 7v1h1v2h1v2h-1v2h-1v1h-1V7h1M8 8V7h1V6h1V5h1V4h1V3h1v16h-1v-1h-1v-1h-1v-1H9v-1H8v-1H4V8h4m-2 2v2h3v1h1v1h1V8h-1v1H9v1H6Z"/>`),
		g.Group(children),
	)
}