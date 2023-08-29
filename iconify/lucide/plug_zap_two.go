package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlugZapTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 2l-2 2.5h3L12 7m-2 7v-3m4 3v-3m-3 8c-1.7 0-3-1.3-3-3v-2h8v2c0 1.7-1.3 3-3 3Zm1 3v-3"/>`),
		g.Group(children),
	)
}