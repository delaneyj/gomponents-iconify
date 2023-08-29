package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 4v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H2v-2h1v-1h1v-1h1v-1h1v-1h1v-2H6V9H5V8H4V7H3V6H2V4h13m-1 12v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H6v1h1v1h1v1h1v1h1v2H9v1H8v1H7v1H6v1h8Z"/>`),
		g.Group(children),
	)
}