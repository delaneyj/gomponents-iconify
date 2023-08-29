package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightThinCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.03 12c0-4.41-3.62-8.03-8.03-8.03c-4.41 0-8.03 3.62-8.03 8.03c0 4.41 3.62 8.03 8.03 8.03c4.41 0 8.03-3.62 8.03-8.03M22 12c0 5.54-4.46 10-10 10S2 17.54 2 12S6.46 2 12 2s10 4.46 10 10m-8.46 1v3l3.96-4l-3.96-4v3H6.5v2"/>`),
		g.Group(children),
	)
}