package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gif(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8H6a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2v-4H7m5-4v8m4-4h3m1-4h-4v8"/>`),
		g.Group(children),
	)
}