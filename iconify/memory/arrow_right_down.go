package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 19h2v-1h1v-1h1v-1h1v-1h1v-2h-2v1h-1v1h-1v-4h-1V9h-1V8h-2V7H3v2h6v1h2v2h1v3h-1v-1h-1v-1H8v2h1v1h1v1h1v1h1"/>`),
		g.Group(children),
	)
}