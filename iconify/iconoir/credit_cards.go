package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCards(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 11.429V18a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-1.5m17-5.071V10a2 2 0 0 0-2-2h-1m3 3.429h-3"/><path d="M19 8v6.5a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h13a2 2 0 0 1 2 2V8Zm0 0H5.5"/></g>`),
		g.Group(children),
	)
}