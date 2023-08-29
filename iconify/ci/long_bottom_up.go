package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongBottomUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.979 8.45l.007 3.55H19V5h-7v2.014l3.55.007L5 17.571L6.429 19l10.55-10.55Z"/>`),
		g.Group(children),
	)
}