package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseMonthlyAmountButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M13.7 12h44.7a1.685 1.685 0 0 1 1.7 1.7v44.7a1.685 1.685 0 0 1-1.7 1.7H13.7a1.685 1.685 0 0 1-1.7-1.7V13.7a1.627 1.627 0 0 1 1.7-1.7Z"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-miterlimit="10" d="M13.7 12h44.7a1.685 1.685 0 0 1 1.7 1.7v44.7a1.685 1.685 0 0 1-1.7 1.7H13.7a1.685 1.685 0 0 1-1.7-1.7V13.7a1.627 1.627 0 0 1 1.7-1.7Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M25 39h22m-22-9h22m-22-9h22m-22 0v18c0 5-1 9-4 12m26-30v28a2.006 2.006 0 0 1-2 2h-6"/></g>`),
		g.Group(children),
	)
}