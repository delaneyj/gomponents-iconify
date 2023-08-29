package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 8h8a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-4m-4 0H4a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h4m4 4v9"/><path d="M19 12v3m0 4a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-7m2.5-4a2.5 2.5 0 0 1-2.457-2.963m2.023-2C7.206 3.014 7.352 3 7.5 3c1.974-.034 3.76 1.95 4.5 5c.74-3.05 2.526-5.034 4.5-5a2.5 2.5 0 1 1 0 5M3 3l18 18"/></g>`),
		g.Group(children),
	)
}