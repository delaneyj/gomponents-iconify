package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SummaryKpi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7.667 27.503L2 22.448l1.331-1.492l5.641 5.031l7.69-7.445a1.928 1.928 0 0 1 2.674-.008l3.624 3.464l5.58-5.973L30 17.39l-5.581 5.975a1.996 1.996 0 0 1-2.838.08l-3.577-3.419l-7.666 7.42a1.963 1.963 0 0 1-2.671.056zM30 11h-4l2-3l2 3zm-8-7h-4v2h4v2h-3v2h3v2h-4v2h4a2.003 2.003 0 0 0 2-2V6a2.002 2.002 0 0 0-2-2zm-6 10h-6v-4a2.002 2.002 0 0 1 2-2h2V6h-4V4h4a2.002 2.002 0 0 1 2 2v2a2.002 2.002 0 0 1-2 2h-2v2h4zM6 12V4H4v1H2v2h2v5H2v2h6v-2H6z"/>`),
		g.Group(children),
	)
}