package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutBottomFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 16v4a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-4h20ZM21 3a1 1 0 0 1 1 1v10H2V4a1 1 0 0 1 1-1h18Z"/>`),
		g.Group(children),
	)
}