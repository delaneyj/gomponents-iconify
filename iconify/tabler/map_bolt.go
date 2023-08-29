package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapBolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 19l-4-2l-6 3V7l6-3l6 3l6-3v8.5M9 4v13m6-10v7.5m4 1.5l-2 3h4l-2 3"/>`),
		g.Group(children),
	)
}