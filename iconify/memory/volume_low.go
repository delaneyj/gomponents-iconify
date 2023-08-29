package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M6 8h4V7h1V6h1V5h1V4h1V3h1v16h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1H6V8m2 2v2h3v1h1v1h1V8h-1v1h-1v1H8Z"/>`),
		g.Group(children),
	)
}