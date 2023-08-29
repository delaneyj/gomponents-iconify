package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutSixFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 10v11H3a1 1 0 0 1-1-1V10h13Zm7 0v10a1 1 0 0 1-1 1h-4V10h5Zm-1-7a1 1 0 0 1 1 1v4H2V4a1 1 0 0 1 1-1h18Z"/>`),
		g.Group(children),
	)
}