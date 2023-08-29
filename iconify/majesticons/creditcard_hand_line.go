package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditcardHandLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 13V5a2 2 0 0 0-2-2h-2m0 0H7a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7M13 3v10.5M9 7v3"/><path d="M21 22v-7.1a7 7 0 0 0-2.052-4.95L16.998 8v6.587c0 .89-1.077 1.337-1.707.707L13.996 14c-.5-.5-1.701-.8-2.502 0c-.8.8-.5 2 0 2.5l5.503 5.5"/></g>`),
		g.Group(children),
	)
}