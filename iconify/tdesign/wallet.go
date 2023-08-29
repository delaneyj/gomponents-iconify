package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h15a3 3 0 0 1 3 3v1h4v14H1V3Zm2 6v10h18V9H3Zm0-2h14V6a1 1 0 0 0-1-1H3v2Zm13 6h3v2h-3v-2Z"/>`),
		g.Group(children),
	)
}