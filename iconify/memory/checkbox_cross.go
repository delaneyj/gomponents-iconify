package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckboxCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M13 12h1v1h1v2h-2v-1h-1v-1h-2v1H9v1H7v-2h1v-1h1v-2H8V9H7V7h2v1h1v1h2V8h1V7h2v2h-1v1h-1v2m5 7H4v-1H3V4h1V3h14v1h1v14h-1v1M5 5v12h12V5H5Z"/>`),
		g.Group(children),
	)
}