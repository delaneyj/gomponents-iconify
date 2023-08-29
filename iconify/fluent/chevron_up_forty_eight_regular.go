package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUpFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M39.634 31.884a1.25 1.25 0 0 1-1.768 0L24 18.018L10.134 31.884a1.25 1.25 0 0 1-1.768-1.768l14.75-14.75a1.25 1.25 0 0 1 1.768 0l14.75 14.75a1.25 1.25 0 0 1 0 1.768Z"/>`),
		g.Group(children),
	)
}