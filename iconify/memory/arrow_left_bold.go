package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 20h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h2v5h8v8h-8v5m-2-4v-3h8V9h-8V6H9v1H8v1H7v1H6v1H5v2h1v1h1v1h1v1h1v1h1Z"/>`),
		g.Group(children),
	)
}