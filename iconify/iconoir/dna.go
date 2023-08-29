package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dna(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 3c0 5.625 8 9 8 9s8 3.375 8 9"/><path d="M20 3c0 5.625-8 9-8 9s-8 3.375-8 9M8 6h11M8 18h11m-8-9h5.5M11 15h5.5"/></g>`),
		g.Group(children),
	)
}