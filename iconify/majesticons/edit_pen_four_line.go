package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditPenFourLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 5l2.293-2.293a1 1 0 0 1 1.414 0l1.586 1.586a1 1 0 0 1 0 1.414L19 8m-3-3l-5.707 5.707a1 1 0 0 0-.293.707V13a1 1 0 0 0 1 1h1.586a1 1 0 0 0 .707-.293L19 8m-3-3l3 3M6 14H5a2 2 0 0 0-2 2v0a2 2 0 0 0 2 2h14a2 2 0 0 1 2 2v0a2 2 0 0 1-2 2h-4"/>`),
		g.Group(children),
	)
}