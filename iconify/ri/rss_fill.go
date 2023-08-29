package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3c9.941 0 18 8.059 18 18h-3c0-8.284-6.716-15-15-15V3Zm0 7c6.075 0 11 4.925 11 11h-3a8 8 0 0 0-8-8v-3Zm0 7a4 4 0 0 1 4 4H3v-4Z"/>`),
		g.Group(children),
	)
}