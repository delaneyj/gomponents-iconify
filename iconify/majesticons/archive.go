package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v1c0 .364-.097.706-.268 1H2.268A1.99 1.99 0 0 1 2 7V6zm1 4v8a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-8H3zm7 2a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-4z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}