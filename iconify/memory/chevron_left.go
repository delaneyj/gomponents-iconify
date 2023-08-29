package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 16v1h1v1h2v-2h-1v-1h-1v-1h-1v-1h-1v-1h-1v-2h1V9h1V8h1V7h1V6h1V4h-2v1h-1v1h-1v1h-1v1H9v1H8v1H7v2h1v1h1v1h1v1h1v1"/>`),
		g.Group(children),
	)
}