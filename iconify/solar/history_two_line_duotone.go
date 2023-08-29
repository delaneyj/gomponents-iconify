package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoryTwoLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2"/><path stroke-linejoin="round" d="M12 9v4h4" opacity=".5"/><circle cx="12" cy="12" r="10" stroke-dasharray=".5 3.5" opacity=".5"/></g>`),
		g.Group(children),
	)
}