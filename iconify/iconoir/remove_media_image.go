package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveMediaImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m3 16l7-3l4 1.818M16 10a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm.879 11.121L19 19m2.121-2.121L19 19m0 0l-2.121-2.121M19 19l2.121 2.121"/><path d="M13 21H3.6a.6.6 0 0 1-.6-.6V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6V13"/></g>`),
		g.Group(children),
	)
}