package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" stroke-miterlimit="10" d="M29.333 17A20.075 20.075 0 0 0 16 35.929C16 47.014 24.954 56 36 56s20-8.986 20-20.071c0-8.74-5.565-16.174-13.333-18.929"/><path d="M36 11v25"/></g>`),
		g.Group(children),
	)
}