package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 5h14v14H5z" opacity=".3"/><path fill="currentColor" d="M3 3v18h18V3H3zm16 16H5V5h14v14z"/>`),
		g.Group(children),
	)
}