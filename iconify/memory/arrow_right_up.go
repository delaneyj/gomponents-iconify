package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 3h2v1h1v1h1v1h1v1h1v2h-2V8h-1V7h-1v4h-1v2h-1v1h-2v1H3v-2h6v-1h2v-2h1V7h-1v1h-1v1H8V7h1V6h1V5h1V4h1"/>`),
		g.Group(children),
	)
}