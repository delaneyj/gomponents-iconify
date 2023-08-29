package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelevisionStop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18a2 2 0 0 1 2 2v12c0 1.1-.9 2-2 2h-5v2H8v-2H3c-1.1 0-2-.9-2-2V5c0-1.11.89-2 2-2m0 2v12h18V5H3m6 3h6v6H9V8Z"/>`),
		g.Group(children),
	)
}