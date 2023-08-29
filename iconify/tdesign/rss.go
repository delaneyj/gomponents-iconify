package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h1c9.389 0 17 7.611 17 17v1h-2v-1c0-8.284-6.716-15-15-15H3V3Zm0 7h1c5.523 0 10 4.477 10 10v1h-2v-1a8 8 0 0 0-8-8H3v-2Zm0 9a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z"/>`),
		g.Group(children),
	)
}