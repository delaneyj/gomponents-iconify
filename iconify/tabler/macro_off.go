package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MacroOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 15a6 6 0 0 0 11.47 2.467"/><path d="M15.53 15.53A6 6 0 0 0 12 21"/><path d="M12 21a6 6 0 0 0-6-6m6 6V11m-1.134-.13a5.007 5.007 0 0 1-3.734-3.723M7 3l3 2l2-2l2 2l3-2v3a5 5 0 0 1-2.604 4.389M3 3l18 18"/></g>`),
		g.Group(children),
	)
}