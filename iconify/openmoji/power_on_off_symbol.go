package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerOnOffSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" d="M36 24v24"/><path stroke-linejoin="round" stroke-miterlimit="10" d="M36 56c11.046 0 20-8.954 20-20s-8.954-20-20-20s-20 8.954-20 20s8.954 20 20 20Z"/></g>`),
		g.Group(children),
	)
}