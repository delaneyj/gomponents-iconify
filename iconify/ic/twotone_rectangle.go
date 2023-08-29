package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6h16v12H4z" opacity=".3"/><path fill="currentColor" d="M2 4v16h20V4H2zm18 14H4V6h16v12z"/>`),
		g.Group(children),
	)
}