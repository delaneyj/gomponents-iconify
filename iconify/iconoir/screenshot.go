package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screenshot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 21.4v-7.006a.6.6 0 0 1 .6-.6h1.173a.6.6 0 0 0 .504-.275l1.446-2.244a.6.6 0 0 1 .504-.275h3.546a.6.6 0 0 1 .504.275l1.446 2.244a.6.6 0 0 0 .504.275H21.4a.6.6 0 0 1 .6.6V21.4a.6.6 0 0 1-.6.6H10.6a.6.6 0 0 1-.6-.6Z"/><path d="M16 19a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM3 18v3h2.5M3 9.5v5M3 6V3h3m3.5 0h5M18 3h3v2.5m0 4.5V8.5"/></g>`),
		g.Group(children),
	)
}