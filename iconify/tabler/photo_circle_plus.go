package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoCirclePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 8h.01m5.954 4.806A9 9 0 0 0 12 3a9 9 0 0 0-9 9a9 9 0 0 0 9.397 8.991"/><path d="m4 15l4-4c.928-.893 2.072-.893 3 0l4 4"/><path d="m14 14l1-1c.928-.893 2.072-.893 3 0m-2 6.33h6m-3-3v6"/></g>`),
		g.Group(children),
	)
}