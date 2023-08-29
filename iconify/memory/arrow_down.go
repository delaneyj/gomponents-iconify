package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 17h-2v-1H9v-1H8v-1H7v-1H6v-2h2v1h1v1h1V4h2v9h1v-1h1v-1h2v2h-1v1h-1v1h-1v1h-1"/>`),
		g.Group(children),
	)
}