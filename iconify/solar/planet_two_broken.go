package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlanetTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-dasharray="3.5 2.5" d="M17.671 6.225c2.102-.415 3.654-.268 4.158.538c1.011 1.616-2.57 5.271-7.998 8.163c-5.429 2.893-10.649 3.927-11.66 2.31c-.516-.823.163-2.178 1.672-3.69"/><path d="M8 5.07A8 8 0 1 1 5.07 8"/></g>`),
		g.Group(children),
	)
}