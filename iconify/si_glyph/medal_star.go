package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedalStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="m8.007 5.996l1.071 3.02h3.54L9.74 11.318l1.094 3.489l-2.865-2.157l-2.864 2.157l1.094-3.489l-2.85-2.302h3.54l1.118-3.02z"/><path d="m8.007 1.947l2.276 6.039H11.5l.469-6.619S12.151.041 10.768.041H5.252c-1.361 0-1.235 1.305-1.235 1.305l.488 6.641h1.229l2.273-6.04z"/></g>`),
		g.Group(children),
	)
}