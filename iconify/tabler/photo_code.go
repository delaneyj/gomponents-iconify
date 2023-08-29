package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 8h.01M11.5 21H6a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v7"/><path d="m3 16l5-5c.928-.893 2.072-.893 3 0l3 3m0 0l1-1c.928-.893 2.072-.893 3 0m2 8l2-2l-2-2m-3 0l-2 2l2 2"/></g>`),
		g.Group(children),
	)
}