package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 18l-2-1l-6 3V7l6-3l6 3l6-3v9M9 4v13m6-10v8m0 4l2 2l4-4"/>`),
		g.Group(children),
	)
}