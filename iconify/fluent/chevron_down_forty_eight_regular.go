package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDownFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.366 16.116a1.25 1.25 0 0 0 0 1.768l14.75 14.75a1.25 1.25 0 0 0 1.768 0l14.75-14.75a1.25 1.25 0 0 0-1.768-1.768L24 29.982L10.134 16.116a1.25 1.25 0 0 0-1.768 0Z"/>`),
		g.Group(children),
	)
}