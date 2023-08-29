package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundUnfoldMore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 5.83l2.46 2.46a.996.996 0 1 0 1.41-1.41L12.7 3.7a.996.996 0 0 0-1.41 0L8.12 6.88a.996.996 0 1 0 1.41 1.41L12 5.83zm0 12.34l-2.46-2.46a.996.996 0 1 0-1.41 1.41l3.17 3.18c.39.39 1.02.39 1.41 0l3.17-3.17a.996.996 0 1 0-1.41-1.41L12 18.17z"/>`),
		g.Group(children),
	)
}