package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitHorizontalSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.5 8a.5.5 0 0 0 0-1h-13a.5.5 0 0 0 0 1h13ZM3 6h1V3a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v3h1V3a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v3Zm0 6V9h1v3a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V9h1v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/>`),
		g.Group(children),
	)
}