package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 18v-5l4-2l4 2v5l-4 2zm4-7V8m1.332-2.666L12 4l4 2v5m-4 2l.661-.331m2.684-1.341L16 11l4 2v3m-1.334 2.667L16 20l-4-2M3 3l18 18"/>`),
		g.Group(children),
	)
}