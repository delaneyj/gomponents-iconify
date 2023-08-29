package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WadOfMoneyLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M9 19c-2.809 0-4.213 0-5.222-.674a4.002 4.002 0 0 1-1.104-1.104C2 16.213 2 14.81 2 12c0-2.809 0-4.213.674-5.222a4 4 0 0 1 1.104-1.104C4.787 5 6.19 5 9 5h6c2.809 0 4.213 0 5.222.674a4 4 0 0 1 1.104 1.104C22 7.787 22 9.19 22 12c0 2.809 0 4.213-.674 5.222a4.003 4.003 0 0 1-1.104 1.104C19.213 19 17.81 19 15 19H9Z"/><path d="M9 9a3 3 0 1 0 0 6m6-6a3 3 0 1 1 0 6" opacity=".5"/><path stroke-linecap="round" d="M9 5v13.5M15 5v13.5"/></g>`),
		g.Group(children),
	)
}