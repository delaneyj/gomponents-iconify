package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RocketOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9.29 9.275A9.03 9.03 0 0 0 9 10a6 6 0 0 0-5 3a8 8 0 0 1 7 7a6 6 0 0 0 3-5c.241-.085.478-.18.708-.283m2.428-1.61A9 9 0 0 0 20 7a3 3 0 0 0-3-3a9 9 0 0 0-6.107 2.864"/><path d="M7 14a6 6 0 0 0-3 6a6 6 0 0 0 6-3m4-8a1 1 0 1 0 2 0a1 1 0 1 0-2 0M3 3l18 18"/></g>`),
		g.Group(children),
	)
}