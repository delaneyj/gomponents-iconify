package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartGrowth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 21H20V4.5a.5.5 0 0 0-1 0V21h-3V8.5a.5.5 0 0 0-1 0V21h-3v-8.5a.5.5 0 0 0-1 0V21H8v-4.5a.5.5 0 0 0-1 0V21H3V2.5a.5.5 0 0 0-1 0v19a.5.5 0 0 0 .5.5h19a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}