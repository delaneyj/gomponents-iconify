package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackpackOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 6h3a6 6 0 0 1 6 6v3m-.129 3.872A3 3 0 0 1 16 21H8a3 3 0 0 1-3-3v-6a5.99 5.99 0 0 1 2.285-4.712M10 6V5a2 2 0 1 1 4 0v1"/><path d="M9 21v-4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v4M3 3l18 18"/></g>`),
		g.Group(children),
	)
}