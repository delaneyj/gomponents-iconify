package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 18l-2-1l-6 3V7l6-3l6 3l6-3v9M9 4v13m6-10v6.5m5 7.5l2-2l-2-2m-3 0l-2 2l2 2"/>`),
		g.Group(children),
	)
}