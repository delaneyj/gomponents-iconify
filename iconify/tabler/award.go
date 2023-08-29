package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Award(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 9a6 6 0 1 0 12 0A6 6 0 1 0 6 9"/><path d="m12 15l3.4 5.89l1.598-3.233l3.598.232l-3.4-5.889M6.802 12l-3.4 5.89L7 17.657l1.598 3.232l3.4-5.889"/></g>`),
		g.Group(children),
	)
}