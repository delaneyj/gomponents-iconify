package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 2h14v1h1v1h1v14h-1v1h-1v1H4v-1H3v-1H2V4h1V3h1V2m13 3V4H5v1H4v12h1v1h12v-1h1V5h-1M6 8h10v2H6V8m0 4h7v2H6v-2Z"/>`),
		g.Group(children),
	)
}