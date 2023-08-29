package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 8h.01M12.5 21H6a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v6.5"/><path d="m3 16l5-5c.928-.893 2.072-.893 3 0l3.5 3.5"/><path d="m14 14l1-1c.679-.653 1.473-.829 2.214-.526M19 22v-6m3 3l-3-3l-3 3"/></g>`),
		g.Group(children),
	)
}