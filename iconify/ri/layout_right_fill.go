package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-4V3h4Zm-6 18H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h12v18Z"/>`),
		g.Group(children),
	)
}