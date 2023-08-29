package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyDinar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 20.01V20m-8-7l2.386-.9a1 1 0 0 0-.095-1.902l-1.514-.404a1 1 0 0 1-.102-1.9L9 7"/><path d="M3 14v1a3 3 0 0 0 3 3h4.161a3 3 0 0 0 2.983-3.32L12 4m4 13l1 1h2a2 2 0 0 0 1.649-3.131L17.996 11"/></g>`),
		g.Group(children),
	)
}