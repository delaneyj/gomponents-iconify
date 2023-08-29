package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeReplaceFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 13.5v5.152a.6.6 0 0 1-.302.52l-6.4 3.658a.6.6 0 0 1-.596 0l-6.4-3.657A.6.6 0 0 1 5 18.652V13m7 9.5V17m11-9L11 1m2 14L1 8m0 0c3-5 7-2 10-7"/><path d="M13 15c3-5 7-2 10-7"/></g>`),
		g.Group(children),
	)
}