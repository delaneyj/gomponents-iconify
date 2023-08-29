package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PensiveFaceBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2.5"><circle cx="12" cy="12" r="9" stroke-linejoin="round"/><path d="M10 16h4"/><path stroke-linejoin="round" d="M14 10.5c.463.188.97.29 1.5.29s1.037-.102 1.5-.29m-10 0c.463.188.97.29 1.5.29s1.037-.102 1.5-.29"/></g>`),
		g.Group(children),
	)
}