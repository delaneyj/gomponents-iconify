package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 21V10h5v10a1 1 0 0 1-1 1h-4Zm-2 0H4a1 1 0 0 1-1-1V10h11v11Zm7-13H3V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v4Z"/>`),
		g.Group(children),
	)
}