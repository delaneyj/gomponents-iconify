package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundSortVerticalBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12s4.477 10 10 10s10-4.477 10-10Z" opacity=".5"/><path d="M11.445 10.245a.75.75 0 1 0 1.11 1.01L13.75 9.94V16a.75.75 0 0 0 1.5 0V9.94l1.195 1.315a.75.75 0 1 0 1.11-1.01l-2.5-2.75a.75.75 0 0 0-1.11 0l-2.5 2.75Z"/><path d="M7.555 12.745a.75.75 0 1 0-1.11 1.01l2.5 2.75a.75.75 0 0 0 1.11 0l2.5-2.75a.75.75 0 0 0-1.11-1.01L10.25 14.06V8a.75.75 0 1 0-1.5 0v6.06l-1.195-1.315Z"/></g>`),
		g.Group(children),
	)
}