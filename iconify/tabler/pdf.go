package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pdf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 8v8h2a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-2zm-7 4h2a2 2 0 1 0 0-4H3v8m14-4h3m1-4h-4v8"/>`),
		g.Group(children),
	)
}