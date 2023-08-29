package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eggs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 22c-3 0-4.868-2.118-5-5c0-3 2-5 5-5c4 0 8.01 2.5 8 5c0 2.5-4 5-8 5z"/><path d="M8 18c-3.03-.196-5-2.309-5-5.38C3 8.313 5.75 3.995 8.5 4c2.614 0 5.248 3.915 5.5 8"/></g>`),
		g.Group(children),
	)
}