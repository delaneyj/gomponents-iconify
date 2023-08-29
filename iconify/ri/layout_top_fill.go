package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutTopFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 10v10a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V10h20Zm-1-7a1 1 0 0 1 1 1v4H2V4a1 1 0 0 1 1-1h18Z"/>`),
		g.Group(children),
	)
}