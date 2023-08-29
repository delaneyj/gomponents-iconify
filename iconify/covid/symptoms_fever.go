package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymptomsFever(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.01 16.607a3.495 3.495 0 1 0 0-6.99a3.495 3.495 0 0 0 0 6.99Zm5.242 6.616a5.243 5.243 0 0 0-10.485 0m20.968-7.821V3.751a2.996 2.996 0 1 0-5.991 0v11.651a4.493 4.493 0 1 0 5.991 0ZM18.74 5.249v11.983"/><path d="M18.74 20.227a1.498 1.498 0 1 0 0-2.995a1.498 1.498 0 0 0 0 2.995ZM2.014 7.745a2.47 2.47 0 0 1 0-3.495a2.472 2.472 0 0 0 0-3.495m3.721 6.99a2.47 2.47 0 0 1 0-3.495a2.472 2.472 0 0 0 0-3.495m3.995 6.99a2.47 2.47 0 0 1 0-3.495a2.473 2.473 0 0 0 0-3.495"/></g>`),
		g.Group(children),
	)
}