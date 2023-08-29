package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 18H9V6l-4 4l-2-1l7-7l7 7l-2 1l-4-4v12z"/>`),
		g.Group(children),
	)
}