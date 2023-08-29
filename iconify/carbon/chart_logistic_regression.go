package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLogisticRegression(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="26" cy="22" r="2" fill="currentColor"/><circle cx="20" cy="22" r="2" fill="currentColor"/><circle cx="14" cy="6" r="2" fill="currentColor"/><circle cx="8" cy="6" r="2" fill="currentColor"/><path fill="currentColor" d="M8.234 24a8.989 8.989 0 0 0 8.538-6.154l2.353-7.06A6.992 6.992 0 0 1 25.766 6H30V4h-4.234a8.989 8.989 0 0 0-8.538 6.154l-2.353 7.06A6.992 6.992 0 0 1 8.234 22H4V2H2v26c0 1.103.897 2 2 2h26v-2H4v-4h4.234Z"/>`),
		g.Group(children),
	)
}