package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSomalia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m9.724 44.022L32 39.097l-9.724 6.926l3.672-11.259l-9.682-6.976l11.99-.034L32 16.522l3.743 11.231l11.99.034l-9.686 6.976l3.677 11.259"/>`),
		g.Group(children),
	)
}