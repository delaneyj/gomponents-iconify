package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckboxMarked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 4h1V3h14v1h-1v1H5v12h12v-6h1v-1h1v8h-1v1H4v-1H3V4m3 5h2v1h1v1h1v1h2v-1h1v-1h1V9h1V8h1V7h1V6h1V5h1V4h2v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6V9Z"/>`),
		g.Group(children),
	)
}