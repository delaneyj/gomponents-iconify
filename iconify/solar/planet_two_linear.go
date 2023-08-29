package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlanetTwoLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M20 12a8 8 0 1 1-16 0a8 8 0 0 1 16 0Z"/><path stroke-dasharray="3.5 2.5" stroke-linecap="round" d="M17.671 6.225c2.102-.415 3.654-.268 4.158.538c1.011 1.616-2.57 5.271-7.998 8.163c-5.429 2.893-10.649 3.927-11.66 2.31c-.516-.823.163-2.178 1.672-3.69"/></g>`),
		g.Group(children),
	)
}