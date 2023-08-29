package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutThreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 10v11H4a1 1 0 0 1-1-1V10h5Zm13 0v10a1 1 0 0 1-1 1H10V10h11Zm-1-7a1 1 0 0 1 1 1v4H3V4a1 1 0 0 1 1-1h16Z"/>`),
		g.Group(children),
	)
}