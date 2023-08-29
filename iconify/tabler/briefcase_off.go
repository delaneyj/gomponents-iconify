package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 7h8a2 2 0 0 1 2 2v8m-1.166 2.818A1.993 1.993 0 0 1 19 20H5a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h2m1.185-2.842A2 2 0 0 1 10 3h4a2 2 0 0 1 2 2v2m-4 5v.01"/><path d="M3 13a20 20 0 0 0 11.905 1.928m3.263-.763A20 20 0 0 0 21 13M3 3l18 18"/></g>`),
		g.Group(children),
	)
}