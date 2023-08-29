package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 8h.01M15 21H6a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v6"/><path d="m3 16l5-5c.928-.893 2.072-.893 3 0l4 4"/><path d="m14 14l1-1c.665-.64 1.44-.821 2.167-.545M19 16v3m0 3v.01"/></g>`),
		g.Group(children),
	)
}