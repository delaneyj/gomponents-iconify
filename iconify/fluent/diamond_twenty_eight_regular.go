package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.732 15.77a2.5 2.5 0 0 1 0-3.536l8.502-8.502a2.5 2.5 0 0 1 3.536 0l8.502 8.502a2.5 2.5 0 0 1 0 3.536l-8.502 8.502a2.5 2.5 0 0 1-3.536 0L3.732 15.77Zm1.06-2.475a1 1 0 0 0 0 1.414l8.503 8.502a1 1 0 0 0 1.414 0l8.502-8.502a1 1 0 0 0 0-1.414L14.71 4.793a1 1 0 0 0-1.414 0l-8.502 8.502Z"/>`),
		g.Group(children),
	)
}