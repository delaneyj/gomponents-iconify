package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RelievedFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M14 15.5a3.981 3.981 0 0 1-2 .535a3.981 3.981 0 0 1-2-.535m4-5c.463.188.97.29 1.5.29s1.037-.102 1.5-.29m-10 0c.463.188.97.29 1.5.29s1.037-.102 1.5-.29"/></g>`),
		g.Group(children),
	)
}