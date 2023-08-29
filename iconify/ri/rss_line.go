package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17a4 4 0 0 1 4 4H3v-4Zm0-7c6.075 0 11 4.925 11 11h-2a9 9 0 0 0-9-9v-2Zm0-7c9.941 0 18 8.059 18 18h-2c0-8.837-7.163-16-16-16V3Z"/>`),
		g.Group(children),
	)
}