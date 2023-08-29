package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.71 12.54a1 1 0 0 0-1.42 0l-3 3A1 1 0 0 0 9.71 17L12 14.66L14.29 17a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm-3-1.08L12 9.16l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-3-3a1 1 0 0 0-1.42 0l-3 3a1 1 0 0 0 1.42 1.42Z"/>`),
		g.Group(children),
	)
}