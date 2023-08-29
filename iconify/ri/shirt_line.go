package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShirtLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.998 20h6v-4h-4v-2h4V6h-2v5l-4-1.6V20Zm-2 0V9.4l-4 1.6V6h-2v14h6Zm-4-16V3h10v1h3a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-16a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h3Zm5 4l3.5-3h-7l3.5 3Z"/>`),
		g.Group(children),
	)
}