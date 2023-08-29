package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 11V9a6 6 0 0 0-6-6v0m6 8v4a6 6 0 0 1-6 6v0a6 6 0 0 1-6-6v-4m12 0h-6m-6 0V9a6 6 0 0 1 6-6v0m-6 8h6m0 0V3"/>`),
		g.Group(children),
	)
}